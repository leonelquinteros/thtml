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
)

const (
	defaultExtensions string = ".thtml .html .js .css"
)

// Service wraps multiple templates within a directory
type Service struct {
	sync.Mutex

	// Filename extensions supported
	exts []string

	// Templates directory
	tplDir string

	// Build input directory
	publicDir string

	// Build output
	buildDir string

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
		return err
	}

	// Init template
	s.tpl = template.New(s.tplDir)

	// Load
	err = filepath.Walk(s.tplDir, s.loadFn)
	if err != nil {
		return err
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
		return err
	}
	content, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	// Create buffer
	buff := new(bytes.Buffer)

	if s.ValidExtension(filepath.Ext(fn)) {
		// Copy template object
		s.Lock()
		tmpTpl, err := s.tpl.Clone()
		s.Unlock()
		if err != nil {
			return err
		}

		// Parse template
		_, err = tmpTpl.New(fn).Parse(string(content))
		if err != nil {
			return err
		}

		// Execute template
		err = tmpTpl.ExecuteTemplate(buff, fn, data)
		if err != nil {
			return err
		}
	} else {
		buff.Write(content)
	}

	// Flush buffer
	_, err = w.Write(buff.Bytes())
	return err
}

// Build compiles all files in the provided directory and outputs the results to the build dir.
// This method is NOT safe to use from multiple/concurrent goroutines
func (s *Service) Build(in, out string) (err error) {
	if s.tpl == nil {
		return NewEmptyTemplateError()
	}

	s.publicDir, err = filepath.Abs(in)
	if err != nil {
		return err
	}
	s.buildDir, err = filepath.Abs(out)
	if err != nil {
		return err
	}

	err = filepath.Walk(s.publicDir, s.buildFn)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) buildFn(filename string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		// Create output
		in := strings.TrimPrefix(filename, s.publicDir)
		out := path.Join(s.buildDir, in)
		f, err := os.OpenFile(out, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		defer f.Close()

		// Render
		return s.Render(f, in, nil)
	}

	return nil
}
