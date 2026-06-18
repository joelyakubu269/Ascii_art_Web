package artgen

import (
	"fmt"
	"strings"
)

func RenderLine(text string, banner map[rune][]string) ([]string, error) {
	result := make([]string, 8)
	for row := 0; row < 8; row++ {
		var builder strings.Builder
		for _, char := range text {
			val, ok := banner[char]
			if ok {
				builder.WriteString(val[row])
			} else {
				return nil, fmt.Errorf("character not found")
			}

		}
		result[row] = builder.String()
	}
	return result, nil
}
