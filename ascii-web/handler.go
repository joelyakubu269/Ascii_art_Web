package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	data := pageData{
		Input: text,
	}

}
