package main

import (
	"asciiartweb/artgen"
	"errors"
	"fmt"
)

var ErrInvalidBanner = errors.New("Banner not allowed")
var ErrBannerNotFound = errors.New("Banner not found")

func Transform(text string, filename string) (string, error) {
	rune, err := artgen.ValidateInput(text)
	if err != nil {
		return "", fmt.Errorf("%c is not a valid character", rune)
	}
	allowed := map[string]bool{
		"standard.txt":   true,
		"shadow.txt":     true,
		"thinkertoy.txt": true,
	}
	if !allowed[filename] {
		return "", ErrInvalidBanner
	}
	fmt.Println("filename received:", filename)

	m, err := artgen.LoadBanner("../ascii-art/" + filename)
	if err != nil {
		return "", ErrBannerNotFound
	}

	return artgen.GenerateArt(text, m)
}
