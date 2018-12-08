# Template HTML

[![GoDoc](https://godoc.org/github.com/leonelquinteros/thtml?status.svg)](https://godoc.org/github.com/leonelquinteros/thtml)
[![GitHub release](https://img.shields.io/github/release/leonelquinteros/thtml.svg)](https://github.com/leonelquinteros/thtml)
[![Build Status](https://travis-ci.org/leonelquinteros/thtml.svg?branch=master)](https://travis-ci.org/leonelquinteros/thtml)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/leonelquinteros/thtml)](https://goreportcard.com/report/github.com/leonelquinteros/thtml)


Static site generator based on Go's text/template with development web server integrated. 


## Install latest version using Go

```
go install github.com/leonelquinteros/thtml
```


## Quick Start

```
$ thtml -h
Usage of thtml:
  -build
    	Build the assets from the -public directory to the -output directory by parsing the -templates directory.
  -exts string
    	Provides a comma separated filename extensions list to support when parsing templates. (default ".html")
  -listen string
    	Run the dev server listening on the provided host:port. (default ":5500")
  -minify
    	Minify the build output. (default false)
  -output string
    	Sets the path for the build output. (default "build")
  -public string
    	Sets the path for the web root. (default "public")
  -run
    	Run a dev web server serving the public directory.
  -templates string
    	Sets the path for the template files. (default "templates")

```


Go template syntax and docs: [https://golang.org/pkg/text/template](https://golang.org/pkg/text/template)


## Creating static websites

### 1. Prepare the environment

After installing the `thtml` tool, we'll create a new directory for our website development files: 

```
mkdir mywebsite
cd mywebsite
``` 

Inside the newly created directory we'll create 2 new directories. 

`templates` will contain all our template files that will be reused from our website pages. 

`public` will contain the website pages and structure.

```
mkdir templates
mkdir public
```


### 2. Create the layout

We'll start by putting together a basic website layout based on [Bootstrap](https://getbootstrap.com/) that will be used by all the pages of our example website.

For that we'll create a new directory for the layouts inside the `templates` directory: 

```
mkdir templates/layouts
```

And then create a new file `templates/layouts/default.html` with the following content: 

```html
<!doctype html>
<html>
    <head>
        <title>Website title</title>

        <!-- Required meta tags for Bootstrap -->
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

        <!-- Bootstrap CSS -->
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/css/bootstrap.min.css" integrity="sha384-PsH8R72JQ3SOdhVi3uxftmaW6Vc51MKb0q5P2rRUpPvrszuE4W1povHYgTpBfshb" crossorigin="anonymous">
    </head>
    <body>
        <div class="container">
            <nav class="navbar navbar-expand-md navbar-light bg-light rounded mb-3">
                <a class="navbar-brand" href="/">THTML</a>
                <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#main-menu" aria-controls="main-menu" aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>

                <div class="collapse navbar-collapse" id="main-menu">
                    <ul class="navbar-nav mr-auto">
                        <li class="nav-item active">
                            <a class="nav-link" href="/">Home</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="/page.html">Link</a>
                        </li>
                    </ul>
                </div>
            </nav>

            <div class="row">
                <div class="col-md-12">
                    {{ block "layout-content" . }}{{ end }}
                </div>
            </div>
        </div>
    </body>
</html>
```

The entire file is a basic Bootstrap template, the only interesting line here is the one that says `{{ block "layout-content" . }}{{ end }}`, which defines a content block that will be extended by our pages. 


### 3. Create the first page

Now lets create a new page at `public/index.html` with the following content: 

```html
{{ template "layouts/default.html" }}

{{ define "layout-content" }}

    <h1>This is the home page</h1>

    <p>
    And this is some content
    </p>

{{ end }}
```

Here we see 2 new interesting lines. First one is `{{ template "layouts/default.html" }}` that loads the template located in `templates/layouts/default.html` by referencing it using a relative path from the `templates` directory. 

Then there is the `{{ define "layout-content" }}` that defines a block of code with the name of `layout-content` that's the name of the content block we defined in our layout file. 


### 4. Run development server

While we create our pages, we need to quickly see what's happening and how they look. For that purpose, we'll use the `run` mode of the `thtml` tool to run a local development web server to serve our website before being compiled to a static form: 

```
thtml -run
```

With all the default options, the tool will use the `templates` and `public` directories properly. After running the command, we can open http://localhost:5500 in our browser to see our home page compiled and running. 

After making any changes to the page or the layout, we can refresh the browser to see these changes while the web server keeps running. Try it! 


### 5. Continue working

Go ahead and create more pages. You can also take the menu to a different file into a different directory (under the `templates` directory). Feel free to get ideas from the [_example](_example) website.


### 6. Build your static website

After you finish developing and your website is ready to go live, you can compile it by running: 

```
thtml -build
```

This will create a static version of your website into the `build` directory by default, but you can configure the output to compile to any directory you want. 

Now you can deploy the contents of the `build` directory to your web server root.  


## Full documentation

[https://godoc.org/github.com/leonelquinteros/thtml](https://godoc.org/github.com/leonelquinteros/thtml)


## Using the thtml/templates package
[https://godoc.org/github.com/leonelquinteros/thtml/templates](https://godoc.org/github.com/leonelquinteros/thtml/templates)

