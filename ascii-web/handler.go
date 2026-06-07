package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func handler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	result, err := transformText(text, banner)

	data := pageData{
		Input: text,
	}

}
