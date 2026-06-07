package main

import (
	"fmt"
)

func transformText(text, bannerFile string) (string, error) {
	r, err := ValidateInput(text)
	if err != nil {
		return "", fmt.Errorf("%c is not a valid ascii character", r)
	}

	allowed := map[string]bool{
		"standard.txt":   true,
		"shadow.txt":     true,
		"thinkertoy.txt": true,
	}
	if !allowed[bannerFile] {
		return "", fmt.Errorf("invalid banner selected")
	}
	charMap, err := LoadBanner(bannerFile)
	if err != nil {
		return "", err
	}

	return GenerateArt(text, charMap), nil
}
