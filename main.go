package main

import (
	"flag"
	"fmt"
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
)

func init() {
	// Parse config flags
	flag.BoolVar(&_run, "run", false, "Run a dev web server serving the public directory. Default: false")
	flag.BoolVar(&_build, "build", false, "Build the assets from the -public directory to the -output directory. Default: false")
	flag.BoolVar(&_minify, "minify", true, "Minify the build output. Default: true")

	flag.StringVar(&_publicPath, "public", "public", "Sets the path for the web root. Default: 'public'")
	flag.StringVar(&_templatesPath, "templates", "templates", "Sets the path for the template files. Default: 'templates'")
	flag.StringVar(&_httpListen, "listen", ":5500", "Run the dev server listening on the provided host:port. Default: ':5500'")
	flag.StringVar(&_outputPath, "output", "build", "Sets the path for the build output. Default: 'build'")
}

func main() {
	flag.Parse()

	if !_run && !_build {
		usage()
		return
	}

	// Build first
	if _build {
		fmt.Println("Coming soon...")
	}

	// Run
	if _run {
		runServer()
	}
}

func usage() {
	fmt.Println("USAGE:")
	fmt.Println("  thtml [OPTIONS] -run")
	fmt.Println("  thtml [OPTIONS] -build")
	fmt.Println("")
	fmt.Println("OPTIONS:")
	fmt.Println("    -public /path/to/public/dir")
	fmt.Println("    -templates /path/to/templates/dir")
	fmt.Println("    -output /path/to/build/output/dir")
	fmt.Println("    -listen ':8080'")
	fmt.Println("    -minify False")
	fmt.Println("")
}
