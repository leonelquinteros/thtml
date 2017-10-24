package templates

import (
	"bytes"
	"sync"
	"testing"

	"github.com/leonelquinteros/gorand"
)

func TestExtensions(t *testing.T) {
	s := new(Service)
	s.AddExtension(".html")
	if len(s.exts) != 1 {
		t.Error("len(s.exts) != 1")
	}

	s.AddExtension(".html")
	if len(s.exts) != 1 {
		t.Error("len(s.exts) != 1")
	}

	s.AddExtension(".js")
	if len(s.exts) != 2 {
		t.Error("len(s.exts) != 2")
	}

	s.RemoveExtension(".txt")
	if len(s.exts) != 2 {
		t.Error("len(s.exts) != 2")
	}

	s.RemoveExtension(".html")
	if len(s.exts) != 1 {
		t.Error("len(s.exts) != 1")
	}

	s.RemoveExtension(".js")
	if len(s.exts) != 0 {
		t.Error("len(s.exts) != 0")
	}
}

func TestLoad(t *testing.T) {
	// Load service
	s := new(Service)
	s.AddExtension(".html")

	err := s.Load("../_example/templates")
	if err != nil {
		t.Fatal(err)
	}

	loaded := s.tpl.Templates()
	expected := []string{
		"components/nav.html",
		"layouts/default.html",
		"views/primary.html",
	}

	for _, e := range expected {
		found := false
		for _, l := range loaded {
			if e == l.Name() {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected template %s to be loaded, but not found", e)
		}
	}
}

func TestRender(t *testing.T) {

}

func TestRace(t *testing.T) {
	wg := new(sync.WaitGroup)

	// Load service
	s := new(Service)
	s.AddExtension(".html")
	err := s.Load("../_example/templates")
	if err != nil {
		t.Error(err)
	}

	// Run basic operations in 1000 goroutines
	for i := 0; i < 1000; i++ {
		ext, err := gorand.GetAlphaString(3)
		if err != nil {
			t.Error(err)
		}
		ext = "." + ext

		wg.Add(1)
		go func(s *Service) {
			s.AddExtension(ext)
			s.ValidExtension(ext)
			s.RemoveExtension(ext)

			buff := &bytes.Buffer{}
			err = s.Render(buff, "../_example/public/index.html", nil)
			if err != nil {
				t.Error(err)
			}
			buff.Reset()

			wg.Done()
		}(s)
	}

	wg.Wait()
}
