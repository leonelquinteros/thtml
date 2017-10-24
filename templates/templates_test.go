package templates

import (
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

func TestExtensionsRace(t *testing.T) {
	wg := new(sync.WaitGroup)

	s := new(Service)
	for i := 0; i < 10000; i++ {
		ext, err := gorand.GetAlphaString(3)
		if err != nil {
			t.Error(err)
		}
		ext = "." + ext

		wg.Add(1)
		go func() {
			s.AddExtension(ext)
			s.ValidExtension(ext)
			s.RemoveExtension(ext)

			wg.Done()
		}()
	}

	wg.Wait()
}
