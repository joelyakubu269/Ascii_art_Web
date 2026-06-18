package main

import (
	"errors"
	"html/template"
	"net/http"
)

var templ = template.Must(template.ParseFiles("templates/index.html"))

func handlers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templ.Execute(w, nil)
		return
	}
	if r.Method != http.MethodPost {
		RenderErrors(w, http.StatusMethodNotAllowed, "method is not allowed")
		return
	}
	str := r.FormValue("text")
	banner := r.FormValue("banner")
	val, err := Transform(str, banner)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidBanner):
			RenderErrors(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, ErrBannerNotFound):
			RenderErrors(w, http.StatusNotFound, err.Error())
		default:
			RenderErrors(w, http.StatusInternalServerError, "interval server error")
		}
		return
	}
	err = templ.Execute(w, pageData{
		Result: val,
	})
	if err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
	}

}

func RenderErrors(w http.ResponseWriter, Status int, msg string) {
	w.WriteHeader(Status)
	err := templ.Execute(w, pageData{
		Error: msg,
	})
	if err != nil {
		http.Error(w, "template rendering failed", http.StatusInternalServerError)
	}
}
