package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . <text> [banner file]")
		os.Exit(1)
	}
	rune, err := ValidateInput(os.Args[1])
	if err != nil {
		fmt.Printf("%c is not a valid ascii character\n", rune)
		os.Exit(1)
	}
	input := os.Args[1]
	banner := "standard.txt"
	if len(os.Args) == 3 {
		banner = os.Args[2]
	}

	charMap, err := LoadBanner(banner)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(GenerateArt(input, charMap))
}
