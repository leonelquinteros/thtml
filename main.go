// thtml is a static website generator based on text/template package.
//
// Includes a development webserver to help creating HTML websites and components
// compiling the templates on the fight, allowing a edit-save-refresh development process.
//
package main

import (
	"flag"
	"fmt"
	"os"
)

// Configuration options
var (
	// -public /path/to/public/dir
	_publicPath string

	// -templates /path/to/templates/dir
	_templatesPath string

	// -build
	_build bool

	// -output
	_outputPath string

	// -run
	_run bool

	// -listen ":5500"
	_httpListen string

	// -minify
	_minify bool

	// -exts .html,.css,.js
	_exts string
)

func init() {
	// Parse config flags
	flag.BoolVar(&_run, "run", false, "Run a dev web server serving the public directory.")
	flag.BoolVar(&_build, "build", false, "Build the assets from the -public directory to the -output directory by parsing the -templates directory.")
	flag.BoolVar(&_minify, "minify", true, "Minify the build output.")
	flag.StringVar(&_publicPath, "public", "public", "Sets the path for the web root.")
	flag.StringVar(&_templatesPath, "templates", "templates", "Sets the path for the template files.")
	flag.StringVar(&_httpListen, "listen", ":5500", "Run the dev server listening on the provided host:port.")
	flag.StringVar(&_outputPath, "output", "build", "Sets the path for the build output.")
	flag.StringVar(&_exts, "exts", ".thtml,.html,.css,.js", "Provides a comma separated filename extensions list to support when parsing templates.")
}

func main() {
	flag.Parse()

	if !_run && !_build {
		fmt.Println("")
		fmt.Println("Run:")
		fmt.Println("     ", os.Args[0], "-h")
		fmt.Println("")
		fmt.Println("To view help")
		fmt.Println("")
		return
	}

	// Build first
	if _build {
		build()
	}

	// Run
	if _run {
		runServer()
	}
}
