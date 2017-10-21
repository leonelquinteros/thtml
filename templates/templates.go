package templates

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
)

// Service ...
type Service struct {
	sync.Mutex

	dir string
	tpl *template.Template
}

// Load ...
func Load(dir string) (*Service, error) {
	s := new(Service)
	err := s.Load(dir)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Load ...
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

	err = filepath.Walk(s.dir, s.walkFn)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) walkFn(path string, info os.FileInfo, err error) error {
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

// Render ...
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
		return errors.New("Template not loaded")
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
