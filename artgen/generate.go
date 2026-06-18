package artgen

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) (string, error) {
	if input == "" {
		return "", nil
	}
	parts := split(input)

	var result strings.Builder
	for i, part := range parts {
		if part == "" {
			if i < len(parts)-1 {
				result.WriteString("\n")
			}
		} else {
			n, err := RenderLine(part, banner)
			if err != nil {
				return "", err
			}
			for _, r := range n {
				result.WriteString(r)
				result.WriteString("\n")
			}
		}
	}
	return result.String(), nil
}
