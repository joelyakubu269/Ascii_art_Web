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
			renderError(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, ErrBannerNotFound):
			renderError(w, http.StatusNotFound, err.Error())
		default:
			renderError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	err = tmpl.ExecuteTemplate(w, "result.html", pageData{
		Input:  text,
		Result: result,
	})
	if err != nil {
		http.Error(w, "template rendering failed", http.StatusInternalServerError)
	}
}
func renderError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	err := tmpl.ExecuteTemplate(w, "errors.html", pageData{
		Error: msg,
	})
	if err != nil {
		http.Error(w, "template rendering failed", http.StatusInternalServerError)
	}
}
