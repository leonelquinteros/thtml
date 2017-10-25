// THTML is a static website generator based on text/template package.
//
// Includes a development webserver to help creating HTML websites and components
// compiling the templates on the fight, allowing a edit-save-refresh development process.
//
// Usage:
//		thtml [OPTIONS] [COMMAND]
//
// [COMMAND] can be the following:
//
//      -build
//	        Build the assets from the [-public] directory to the [-output] directory by parsing the [-templates] directory.
//
//      -run
//	        Run development webserver listening to [-listen] to build pages on-the-fly.
//
// [OPTIONS] are:
//
//      -exts string
// 	        Provides a comma separated filename extensions list to support when parsing templates. (default ".thtml,.html,.css,.js")
//
//      -listen string
// 	        Run the dev server listening on the provided host:port. (default ":5500")
//
//      -minify
// 	        Minify the build output. (default true)
//
//      -output string
// 	        Sets the path for the build output. (default "build")
//
//      -public string
// 	        Sets the path for the web root. (default "public")
//
//      - templates string
// 	        Sets the path for the template files. (default "templates")
//
//
//
// Getting started
//
// By running `thtml -run` on a directory, the tool will use the default options that assume the following directory structure:
//
//		./public
//		./templates
//		./build
//
// The `public` directory and all its sub-directories will contain the source files for the template website structure,
// including assets that have to be included during the static website compilation.
// It works as the "web root" directory of the development webserver and contains the desired final structure.
//
// The `templates` directory and all its sub-directories is where all the reusable templates have to be placed
// so they can be used from the pages in the `public` directory and from other templates.
//
// The `build` directory will be used to write the output of the static website compilation and then deploy it.
//
// Check the `_example` directory on the repository to see a simple layout: https://github.com/leonelquinteros/thtml/tree/master/_example
//
// After creating the website using the development webserver, it can be built running `thtml -build`
// and the content of the `build` directory can be deployed to any static web server on production.
//
//
//
// Template syntax
//
// THTML uses pure Go's text/template package to render the templates together.
// Check the package documentation for details about the syntax: https://golang.org/pkg/text/template/
//
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
