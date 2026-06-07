package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	ValidateInput(text)
	banner := r.FormValue("banner")
	allowed := map[string]bool{
		"standard.txt":   true,
		"shadow.txt":     true,
		"thinkertoy.txt": true,
	}
	if allowed[banner] {
		LoadBanner(banner)
	}
	data := pageData{
		Input: text,
	}

}
