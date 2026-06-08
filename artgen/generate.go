package artgen

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}

	parts := SplitInput(input)

	var result strings.Builder

	for i, part := range parts {
		if part == "" {
			if i < len(parts)-1 {
				result.WriteString("\n")

			} else if part == "" && i == len(parts)-1 && parts[i-1] != "" {
				result.WriteString(strings.Repeat("\n", 1)) // this line is used to reduce repetition and improve professionalism.
				/*
					just like this:
					result.WriteString("\n")
					result.WriteString("\n")
					result.WriteString("\n")
					result.WriteString("\n")
					result.WriteString("\n")
					result.WriteString("\n")
					result.WriteString("\n")
					result.WriteString("\n")

					so instead of writing it this way, you can just write it on a single line with less code just like how it is used here.
				*/
			}

		} else {
			rows := RenderLine(part, banner)
			for _, row := range rows {
				result.WriteString(row)
				result.WriteString("\n")
			}
		}
	}
	return result.String()
}
