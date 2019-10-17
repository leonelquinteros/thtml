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
  -init
    	Creates a new project structure into the current directory.
  -listen string
    	Run the dev server listening on the provided host:port. (default "localhost:5500")
  -minify
    	Minify the build output. (default true)
  -output string
    	Sets the path for the build output. (default "build")
  -public string
    	Sets the path for the web root. (default "public")
  -run
    	Run a dev web server serving the public directory.
  -templates string
    	Sets the path for the template files. (default "templates")
  -version
    	Prints version number.

```


Go template syntax and docs: [https://golang.org/pkg/text/template](https://golang.org/pkg/text/template)


## Creating static websites

### 1. Prepare the environment

After installing the `thtml` tool, we'll create a new project directory for our website development files. 
thtml will create a scaffolded website structure, based on a dummy Twitter Bootstrap template, were we can start working on. 

```
$ mkdir mywebsite
$ cd mywebsite
$ thtml -init
2019/10/16 22:43:23 => Creating directory structure...
2019/10/16 22:43:23 public/css
2019/10/16 22:43:23 public/js
2019/10/16 22:43:23 public/img
2019/10/16 22:43:23 templates/layouts
2019/10/16 22:43:23 templates/components
2019/10/16 22:43:23 => Fetching files...
2019/10/16 22:43:23 public/css/bootstrap.min.css
2019/10/16 22:43:24 public/css/bootstrap.min.css.map
2019/10/16 22:43:25 public/img/bg-banner.jpg
2019/10/16 22:43:25 public/js/bootstrap.min.js
2019/10/16 22:43:25 public/js/bootstrap.min.js.map
2019/10/16 22:43:25 public/js/jquery-3.2.1.slim.min.js
2019/10/16 22:43:26 public/js/popper.min.js
2019/10/16 22:43:26 public/index.html
2019/10/16 22:43:26 public/page.html
2019/10/16 22:43:27 templates/components/footer.html
2019/10/16 22:43:27 templates/components/head.html
2019/10/16 22:43:27 templates/components/nav.html
2019/10/16 22:43:27 templates/layouts/default.html
2019/10/16 22:43:27 templates/layouts/two-columns.html
Finished.
$
``` 

Inside the project directory 2 new directories are created. 

`templates` will contain all our template files that will be reused from our website pages. 

`public` will contain the website pages and structure.


### 2. Work with the layout

We'll start by putting together a basic website layout based on [Bootstrap](https://getbootstrap.com/) that will be used by all the pages of our example website.

The newly created file `templates/layouts/default.html` has the following content: 

```html
{{ template "components/head.html" }}

<div class="container">
    <div class="row">
        <div class="col-md-12">
            <small>
                Edit <code>templates/layouts/default.html</code> to change this layout
            </small>
        </div>
    </div>
    {{ template "components/nav.html" }}

    <div class="row">
        <div class="col-md-12">
            {{ block "view-content" . }}{{ end }}
        </div>
    </div>
</div>

{{ template "components/footer.html" }}
```

The entire file is a basic Bootstrap template, the only interesting line here is the one that says `{{ block "view-content" . }}{{ end }}`, which defines a content block that will be extended by our pages. 


### 3. Create the first page

Now lets see the file at `public/index.html` with the following content: 

```html
{{ template "layouts/default.html" }}

{{ define "layout-title" }}Home title{{ end }}

{{ define "view-content" }}

    <h1>Home page</h1>

    <p>
        <img class="img-fluid" src="/img/bg-banner.jpg" alt="Image" />
        <br />
        
        <small>
            Edit <code>public/index.html</code> to change this page
        </small>
    </p>

    <p>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
        Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. 
        Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. 
        Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
    </p>

    <p>
        <a href="/page.html" class="btn btn-primary">I'm a button!</a>
    </p>

{{ end }}
```

Here we see 2 new interesting lines. First one is `{{ template "layouts/default.html" }}` that loads the template located in `templates/layouts/default.html` by referencing it using a relative path from the `templates` directory. 

Then there is the `{{ define "view-content" }}` that defines a block of code with the name of `view-content` that's the name of the content block we defined in our layout file. 


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

