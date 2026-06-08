package main

import (
	"asciiartweb/artgen"
	"errors"
	"fmt"
)

var ErrInvalidBanner = errors.New("invalid banner selected")
var ErrBannerNotFound = errors.New("banner not found")

func transformText(text, bannerFile string) (string, error) {
	r, err := artgen.ValidateInput(text)
	if err != nil {
		return "", fmt.Errorf("%c is not a valid ascii character", r)
	}

	allowed := map[string]bool{
		"standard.txt":   true,
		"shadow.txt":     true,
		"thinkertoy.txt": true,
	}
	if !allowed[bannerFile] {
		return "", ErrInvalidBanner
	}
	charMap, err := artgen.LoadBanner(bannerFile)
	if err != nil {
		return "", ErrBannerNotFound
	}

	return artgen.GenerateArt(text, charMap), nil
}
