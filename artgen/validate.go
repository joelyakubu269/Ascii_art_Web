package artgen

import (
	"fmt"
)

func ValidateInput(text string) (rune, error) {

	for _, char := range text {
		if char != ' ' && char < 32 || char > 126 {
			return char, fmt.Errorf("invalid character in text input")
		}
	}
	return 0, nil
}
