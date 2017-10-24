# Template HTML

[![GoDoc](https://godoc.org/github.com/leonelquinteros/thtml?status.svg)](https://godoc.org/github.com/leonelquinteros/thtml)
[![GitHub release](https://img.shields.io/github/release/leonelquinteros/thtml.svg)](https://github.com/leonelquinteros/thtml)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/leonelquinteros/thtml)](https://goreportcard.com/report/github.com/leonelquinteros/thtml)


Static site generator based on Go's text/template with development web server integrated. 


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


## Using the thtml/templates package

[![GoDoc](https://godoc.org/github.com/leonelquinteros/thtml/templates?status.svg)](https://godoc.org/github.com/leonelquinteros/thtml/templates)

...