package main

import (
	"fmt"
	"net/http"
)

func maian() {
	http.HandleFunc("GET /ascii-art", handler)
	http.HandleFunc("POST/ascii-art", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
