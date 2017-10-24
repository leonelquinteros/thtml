package main

import (
	"log"
	"strings"

	"github.com/leonelquinteros/thtml/templates"
)

func build() {
	// Load comma separated extensions list
	exts := strings.Split(_exts, ",")

	// Load templates
	tpl, err := templates.Load(_templatesPath, exts...)
	if err != nil {
		log.Fatalf("Error loading templates from '%s': %s", _templatesPath, err)
	}

	// Build
	err = tpl.Build(_publicPath, _outputPath)
	if err != nil {
		log.Fatalf("Error compiling templates from '%s' to '%s': %s", _publicPath, _outputPath, err)
	}
}
