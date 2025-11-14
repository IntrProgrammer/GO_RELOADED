package main

import (
	"GO_RELOADED/formatter"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go-reloaded <input> <output>")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	content, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	f := formatter.New()
	result := f.Format(string(content))

	err = os.WriteFile(outputPath, []byte(result), 0644)
	if err != nil {
		log.Fatalf("Error writing output: %v", err)
	}

	fmt.Printf("Successfully processed %s -> %s\n", inputPath, outputPath)
}
