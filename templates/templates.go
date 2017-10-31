// Package templates compiles all templates in a directory tree into a Service
// that provides a Render method to execute other templates inside its context.
//
// Example
//
// The following example program will load a template directory tree defined as a constant
// and then render a template file from another directory using the loaded Service into the standard output.
//
//  package main
//
//  import (
//      "os"
//      "path"
//      "github.com/leonelquinteros/thtml/templates"
//  )
//
//  const (
//      _templates = "/path/to/templates"
//      _public    = "/path/to/web/root"
//  )
//
//  func main() {
//      tplService, err := templates.Load(_templates)
//      if err != nil {
//          panic(err.Error())
//      }
//
//      tplService.Render(os.Stdout, path.Join(_public, "index.html"), nil)
//  }
//
package templates

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
)

const (
	defaultExtensions string = ".html"
)

// Service is the template handler.
// After Load()'ing a directory tree, will be ready to render any template into its context.
type Service struct {
	sync.Mutex

	// Filename extensions supported
	exts []string

	// Templates directory
	tplDir string

	// Build input directory
	publicDir string

	// Build output directory
	buildDir string

	// Minify output
	minify bool

	// Template wrapper
	tpl *template.Template
}

// Load creates a new *templates.Service object and loads the templates in the provided directory.
// Custom set of filename extensions can be supplied
func Load(dir string, extensions ...string) (*Service, error) {
	s := new(Service)

	if len(extensions) == 0 {
		extensions = strings.Split(defaultExtensions, " ")
	}
	for _, ext := range extensions {
		s.AddExtension(ext)
	}

	err := s.Load(dir)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Minify sets the configuration to minify the output
func (s *Service) Minify(m bool) {
	s.minify = m
}

// AddExtension adds a new filename extension (i.e. ".txt") to the list of extensions to support.
// Extensions not supported will be rendered and/or compiled as they are without template parsing.
// This method is safe to use from multiple/concurrent goroutines.
func (s *Service) AddExtension(ext string) {
	// Sync
	s.Lock()
	defer s.Unlock()

	// Check
	if s.exts == nil {
		s.exts = make([]string, 0)
	}

	// Avoid repeated
	for i := range s.exts {
		if s.exts[i] == ext {
			return
		}
	}

	// Add
	s.exts = append(s.exts, strings.TrimSpace(ext))
}

// RemoveExtension deletes a supported filename extension from the list.
// This method is safe to use from multiple/concurrent goroutines.
func (s *Service) RemoveExtension(ext string) {
	s.Lock()
	defer s.Unlock()

	for i := range s.exts {
		if s.exts[i] == ext {
			s.exts = append(s.exts[:i], s.exts[i+1:]...)
			return
		}
	}
}

// ValidExtension returns true if the filename extension provided is supported.
// This method is safe to use from multiple/concurrent goroutines.
func (s *Service) ValidExtension(ext string) bool {
	s.Lock()
	defer s.Unlock()

	for _, e := range s.exts {
		if e == ext {
			return true
		}
	}

	return false
}

// Load takes a directory path and loads all templates on it.
// This method is NOT safe to use from multiple/concurrent goroutines
func (s *Service) Load(dir string) error {
	var err error

	// Parse dir name
	s.tplDir, err = filepath.Abs(dir)
	if err != nil {
		return NewError("Error locating template directory " + dir + ": " + err.Error())
	}

	// Init template
	s.tpl = template.New(s.tplDir)

	// Load
	err = filepath.Walk(s.tplDir, s.loadFn)
	if err != nil {
		return NewError("Error loading templates from " + dir + ": " + err.Error())
	}

	return nil
}

func (s *Service) loadFn(path string, info os.FileInfo, err error) error {
	if !info.IsDir() && s.ValidExtension(filepath.Ext(path)) {
		// Load content
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Set tpl name
		n := strings.TrimPrefix(path, s.tplDir)
		for len(n) > 0 && n[0] == '/' {
			n = n[1:]
		}

		// Load template.
		_, err = s.tpl.New(n).Parse(string(content))
		if err != nil {
			return err
		}
	}

	return nil
}

// Render compiles the provided template filename in the loaded templates and writes the output to the provided io.Writer.
// This method is safe to use from multiple/concurrent goroutines
func (s *Service) Render(w io.Writer, filename string, data interface{}) error {
	// Check load
	s.Lock()
	empty := (s.tpl == nil)
	s.Unlock()
	if empty {
		return NewEmptyTemplateError()
	}

	// Load content
	fn, err := filepath.Abs(filename)
	if err != nil {
		return NewError("Error locating template " + filename + ": " + err.Error())
	}
	content, err := ioutil.ReadFile(fn)
	if err != nil {
		return NewError("Error reading template " + filename + ": " + err.Error())
	}

	// Create buffer
	buff := new(bytes.Buffer)

	if s.ValidExtension(filepath.Ext(fn)) {
		// Copy template object
		s.Lock()
		tmpTpl, err := s.tpl.Clone()
		s.Unlock()
		if err != nil {
			return NewError("Error cloning template " + filename + ": " + err.Error())
		}

		// Parse template
		_, err = tmpTpl.New(fn).Parse(string(content))
		if err != nil {
			return NewError("Error parsing template " + filename + ": " + err.Error())
		}

		// Execute template
		err = tmpTpl.ExecuteTemplate(buff, fn, data)
		if err != nil {
			return NewError("Error executing template " + filename + ": " + err.Error())
		}
	} else {
		buff.Write(content)
	}

	// Minifier
	ext := filepath.Ext(fn)
	var mime string
	switch ext {
	case ".js":
		mime = "text/javascript"
	case ".css":
		mime = "text/css"
	case ".html":
		mime = "text/html"
	}

	result := new(bytes.Buffer)
	if mime != "" && s.minify {
		m := minify.New()
		m.AddFunc("text/css", css.Minify)
		m.AddFunc("text/html", html.Minify)
		m.AddFunc("text/javascript", js.Minify)

		// Configure HTML minifier
		m.Add("text/html", &html.Minifier{
			KeepDocumentTags: true,
			KeepEndTags:      true,
		})

		// Minify
		err = m.Minify(mime, result, buff)
		if err != nil {
			// If minify fails, use raw version silently.
			result.Reset()
			result.Write(buff.Bytes())
		}
	} else {
		result.Write(buff.Bytes())
	}

	// Flush buffer
	_, err = w.Write(result.Bytes())
	if err != nil {
		return NewError("Error writing template output " + filename + ": " + err.Error())
	}

	return nil
}

// Build compiles all files in the provided directory and outputs the results to the build dir.
// This method is NOT safe to use from multiple/concurrent goroutines
func (s *Service) Build(in, out string) (err error) {
	if s.tpl == nil {
		return NewEmptyTemplateError()
	}

	s.publicDir, err = filepath.Abs(in)
	if err != nil {
		return NewError("Error locating input directory " + in + ": " + err.Error())
	}
	s.buildDir, err = filepath.Abs(out)
	if err != nil {
		return NewError("Error locating output directory " + in + ": " + err.Error())
	}

	err = filepath.Walk(s.publicDir, s.buildFn)
	if err != nil {
		return NewError("Error building output: " + err.Error())
	}

	return nil
}

func (s *Service) buildFn(filename string, info os.FileInfo, err error) error {
	// Ensure directories
	if !info.IsDir() {
		// Create output
		in := strings.TrimPrefix(filename, s.publicDir)
		out, err := filepath.Abs(path.Join(s.buildDir, in))
		if err != nil {
			return err
		}

		err = os.MkdirAll(path.Dir(out), 0755)
		if err != nil {
			return err
		}

		// Open file
		f, err := os.OpenFile(out, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		defer f.Close()

		// Render

		return s.Render(f, filename, nil)
	}

	return nil
}
