package main

import (
	"flag"
	"fmt"
	"lwc/internal/counter"
	"lwc/pkg/utils"
	"os"
)

func main() {
	filename := flag.String("file", "", "File to process")
	countBytes := flag.Bool("c", false, "Count bytes")
	countWords := flag.Bool("w", false, "Count words")
	countLines := flag.Bool("l", false, "Count lines")
	countChars := flag.Bool("m", false, "Count characters")

	flag.Parse()

	if *filename == "" {
		fmt.Println("Usage: wc -file <filename> [-c] [-w] [-l] [-m]")
		os.Exit(1)
	}

	content, err := utils.ReadFile(*filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bytes, chars, words, lines := counter.Count(content)

	if *countBytes {
		fmt.Println("Byte Count:", bytes)
	}
	if *countWords {
		fmt.Println("Word Count:", words)
	}
	if *countLines {
		fmt.Println("Line Count:", lines)
	}
	if *countChars {
		fmt.Println("Character Count:", chars)
	}

	if !*countBytes && !*countWords && !*countLines && !*countChars {
		fmt.Println("Byte Count:", bytes)
		fmt.Println("Word Count:", words)
		fmt.Println("Line Count:", lines)
		fmt.Println("Character Count:", chars)
	}
}
