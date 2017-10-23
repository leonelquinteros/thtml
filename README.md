# Template HTML

Static site generator with Go's text/template syntax and development webserver integrated. 

## Install

```
go install github.com/leonelquinteros/thtml
```

## Quick Start

```
$ thtml
USAGE:
  thtml [OPTIONS] -run
  thtml [OPTIONS] -build

OPTIONS:
    -public /path/to/public/dir
    -templates /path/to/templates/dir
    -output /path/to/build/output/dir
    -listen ':8080'
    -minify False
```