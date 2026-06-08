package main

import (
	"asciiartweb/artgen"
	"fmt"
	"net/http"
)

func transformText(text, bannerFile string) (string, error) {
	r, err := artgen.ValidateInput(text)
	if err != nil {
		return "", fmt.Errorf("%c is not a valid ascii character", r)
	}
	var ErrInvalidBanner = errors.New("invalid banner selected")
	var ErrBannerNotFound = errors.New("banner not found")
	allowed := map[string]bool{
		"standard.txt":   true,
		"shadow.txt":     true,
		"thinkertoy.txt": true,
	}
	if !allowed[bannerFile] {
		return "",
	}
	charMap, err := artgen.LoadBanner(bannerFile)
	if err != nil {
		return "", err
	}

	return artgen.GenerateArt(text, charMap), nil
}
