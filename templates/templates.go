package templates

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
)

// Service wraps multiple templates within a directory
type Service struct {
	sync.Mutex

	dir string
	tpl *template.Template
}

// Load creates a new *templates.Service object and loads the templates in the provided directory.
func Load(dir string) (*Service, error) {
	s := new(Service)
	err := s.Load(dir)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Load takes a directory path and loads all templates on it.
func (s *Service) Load(dir string) (err error) {
	s.Lock()
	defer s.Unlock()

	// Parse dir
	s.dir, err = filepath.Abs(dir)
	if err != nil {
		return err
	}

	// Init template
	s.tpl = template.New(s.dir)

	err = filepath.Walk(s.dir, s.loadFn)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) loadFn(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		// Load content
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Set tpl name
		n := strings.TrimPrefix(path, s.dir)
		for len(n) > 0 && n[0] == '/' {
			n = n[1:]
		}

		// Load template. Ignore errors
		_, err = s.tpl.New(n).Parse(string(content))
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

// Render compiles the provided template filename in the loaded templates and writes the output to the provided io.Writer.
func (s *Service) Render(w io.Writer, filename string, data interface{}) error {
	fn, err := filepath.Abs(filename)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	if s.tpl == nil {
		return NewEmptyTemplateError()
	}

	// Copy Template
	s.Lock()
	tmp := *s.tpl
	tmpTpl := &tmp
	s.Unlock()

	_, err = tmpTpl.New(fn).Parse(string(content))
	if err != nil {
		return err
	}

	err = tmpTpl.ExecuteTemplate(w, fn, data)
	if err != nil {
		return err
	}

	return nil
}

// Build compiles all files in the provided directory and outputs the results to the build dir.
func (s *Service) Build(dir string) error {
	if s.tpl == nil {
		return NewEmptyTemplateError()
	}

	dn, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	err = filepath.Walk(dn, s.buildFn)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) buildFn(path string, info os.FileInfo, err error) error {
	return nil
}
