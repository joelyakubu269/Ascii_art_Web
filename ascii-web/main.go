package main

import "net/http"

func main() {

	http.HandleFunc("/ascii", handlers)
	http.ListenAndServe(":8080", nil)
}
