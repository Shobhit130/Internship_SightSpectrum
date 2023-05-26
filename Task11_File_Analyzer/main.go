package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	// Prompt the user to enter the file path
	fmt.Print("Enter the path of the text file: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filePath := scanner.Text()

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open the file: %v", err)
	}
	defer file.Close()

	// Initialize counters
	lineCount := 0
	wordCount := 0
	characterCount := 0
	alphabeticCount := 0
	digitCount := 0
	whitespaceCount := 0
	otherCount := 0

	// Create a scanner to read the file line by line
	scanner = bufio.NewScanner(file)
	// Setting the splitting behavior of scanner to split the input into lines. 
	scanner.Split(bufio.ScanLines)
	// iterate over each line of the file and perform further processing
	for scanner.Scan() {
		line := scanner.Text()

		// Increment line count
		lineCount++

		// Count characters
		characterCount += len(line)

		// Split the line into words using custom delimiters
		words := strings.FieldsFunc(line, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r)
		})

		// Count words
		wordCount += len(words)

		// Count alphabetic characters, digit characters, whitespace characters, and other characters
		for _, word := range words {
			for _, char := range word {
				// Count alphabetic characters
				if unicode.IsLetter(char) {
					alphabeticCount++
				} else if unicode.IsDigit(char) {
					// Count digit characters
					digitCount++
				}
			}
		}

		// Count commas, periods and question marks (other characters)
		otherCount += strings.Count(line, ",") + strings.Count(line, ".") + strings.Count(line, "?")

		// Count whitespace characters
		whitespaceCount += strings.Count(line, " ")
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read the file: %v", err)
	}

	// Print the statistics
	fmt.Printf("File Statistics:\n")
	fmt.Printf("Number of lines: %d\n", lineCount)
	fmt.Printf("Number of words: %d\n", wordCount)
	fmt.Printf("Number of characters: %d\n", characterCount)
	fmt.Printf("Number of alphabetic characters: %d\n", alphabeticCount)
	fmt.Printf("Number of digit characters: %d\n", digitCount)
	fmt.Printf("Number of whitespace characters: %d\n", whitespaceCount)
	fmt.Printf("Number of other characters: %d\n", otherCount)
}
