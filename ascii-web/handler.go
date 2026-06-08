package main

import (
	"errors"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func handler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	result, err := transformText(text, banner)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidBanner):
			renderError(w, 400, err.Error())
		}
		tmpl.ExecuteTemplate(w, "error.html", pageData{
			Error: err.Error(),
		})
		return
	}

	tmpl.ExecuteTemplate(w, "result.html", pageData{
		Input:  text,
		Result: result,
	})
}
func renderError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	tmpl.ExecuteTemplate(w, "errors.html", pageData{
		Error: msg,
	})
}
