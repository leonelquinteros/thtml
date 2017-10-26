package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/leonelquinteros/thtml/templates"
)

// Handler
type thtmlHandler struct{}

func (h thtmlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Catch panics
	defer func() {
		if err := recover(); err != nil {
			// Get stack trace
			stack := make([]byte, 1<<16)
			runtime.Stack(stack, false)

			// Log panic
			log.Printf("PANIC! :: %v", err)
			log.Print(string(stack))
		}

		// Log
		log.Printf("%s %s", r.Method, r.URL.String())
	}()

	// Construct path
	p := h.cleanPath(_publicPath + r.URL.EscapedPath())

	// Check if file exists and if it's a file
	if info, err := os.Stat(p); err == nil && !info.IsDir() {
		// Load comma separated extensions list
		exts := strings.Split(_exts, ",")

		// Load templates
		tpl, err := templates.Load(_templatesPath, exts...)
		if err != nil {
			log.Fatalf("Error loading templates from '%s': %s", _templatesPath, err)
		}

		// Render to buffer
		buff := new(bytes.Buffer)
		err = tpl.Render(buff, p, nil)
		if err != nil {
			buff.Write([]byte(err.Error()))
		}
		content := buff.Bytes()

		// Detect content type
		ext := filepath.Ext(p)
		switch ext {
		case ".html":
			w.Header().Set("Content-Type", "text/html; charset=utf-8")

		case ".js":
			w.Header().Set("Content-Type", "application/javascript; charset=utf-8")

		case ".css":
			w.Header().Set("Content-Type", "text/css; charset=utf-8")

		case ".svg":
			w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")

		default:
			w.Header().Set("Content-Type", http.DetectContentType(content))
		}

		// Flush
		w.Write(content)
	} else {
		w.WriteHeader(404)
	}
}

// cleanPath normalizes requeste filenames
func (h thtmlHandler) cleanPath(p string) string {
	p = path.Clean(p)

	// Catch routes without ".html" and dir names without /index.html
	if info, err := os.Stat(p); err != nil || info.IsDir() {
		if info, err := os.Stat(p + ".html"); err == nil && !info.IsDir() {
			p += ".html"
		} else if info, err := os.Stat(p + "index.html"); err == nil && !info.IsDir() {
			p += "index.html"
		} else if info, err := os.Stat(p + "/index.html"); err == nil && !info.IsDir() {
			p += "/index.html"
		}
	}

	return p
}

func runServer() {
	// Routes
	http.Handle("/", thtmlHandler{})

	// Server
	s := &http.Server{
		Addr:           _httpListen,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Starting dev web server on '%s'", _httpListen)

	// Lock
	log.Print(s.ListenAndServe())
}
