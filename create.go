package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const rawHost = "https://raw.githubusercontent.com/leonelquinteros/thtml/master/_example/"

// l logs errors
func l(err error) {
	if err != nil {
		log.Println(err)
	}
}

// create scaffolds a new static web project
func create() {
	dirs := []string{
		"public/css",
		"public/js",
		"public/img",
		"templates/layouts",
		"templates/views",
		"templates/components",
	}

	files := []string{
		"public/css/bootstrap.min.css",
		"public/css/bootstrap.min.css.map",
		"public/img/bg-banner.jpg",
		"public/js/bootstrap.min.js",
		"public/js/bootstrap.min.js.map",
		"public/js/jquery-3.2.1.slim.min.js",
		"public/js/popper.min.js",
		"public/index.html",
		"public/page.html",
		"templates/components/nav.html",
		"templates/layouts/default.html",
		"templates/views/primary.html",
	}

	log.Println("=> Creating directory structure...")
	for _, d := range dirs {
		log.Println(d)
		l(os.MkdirAll(d, 0755))
	}

	log.Println("=> Fetching files...")
	for _, f := range files {
		log.Println(f)

		r, err := http.Get(rawHost + f)
		l(err)

		buff, err := ioutil.ReadAll(r.Body)
		l(err)

		ioutil.WriteFile(f, buff, 0755)
	}

	fmt.Println("Finished.")
}
