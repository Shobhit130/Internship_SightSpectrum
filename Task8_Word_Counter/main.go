package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	// Open the text file
	filePath := "file.txt"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a map to store word counts
	wordCount := make(map[string]int)

	// Total number of words
	totalWords := 0

	// Iterate through each line of the file
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into words
		words := strings.Fields(line)
		// Update word counts
		for _, word := range words {
			wordCount[word]++
			totalWords++
		}
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %s", err)
	}

	// Create a slice to store word-frequency pairs
	wordFreq := make([]WordFrequency, 0, len(wordCount))

	// Convert word counts to word-frequency pairs
	for word, count := range wordCount {
		wordFreq = append(wordFreq, WordFrequency{Word: word, Frequency: count})
	}

	// Sort the word-frequency pairs by frequency in descending order
	sort.Slice(wordFreq, func(i, j int) bool {
		return wordFreq[i].Frequency > wordFreq[j].Frequency
	})

	// Display all words with their counts
	fmt.Println("All words and their counts (indexed and sorted by frequency):")
	for i, wf := range wordFreq {
		fmt.Printf("%d. Word: %s, Frequency: %d\n", i+1, wf.Word, wf.Frequency)
	}

	// Display the total number of words
	fmt.Println("Total number of words:", totalWords)

	// Display the top 10 most frequent words
	fmt.Println("Top 10 most frequent words:")
	for i, wf := range wordFreq[:10] {
		fmt.Printf("%d. Word: %s, Frequency: %d\n", i+1, wf.Word, wf.Frequency)
	}
}

// WordFrequency represents a word and its frequency
type WordFrequency struct {
	Word      string
	Frequency int
}
