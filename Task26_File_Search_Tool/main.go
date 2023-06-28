package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// Get user input for search criteria
	searchDir := getInput("Enter the directory to search: ")
	fileName := getInput("Enter the file name (leave blank for any): ")
	fileType := getInput("Enter the file type (leave blank for any): ")
	modifiedDate := getInput("Enter the modified date (YYYY-MM-DD) (leave blank for any): ")

	// Parse modified date if provided
	var modifiedAfter time.Time
	if modifiedDate != "" {
		parsedDate, err := time.Parse("2006-01-02", modifiedDate)
		if err != nil {
			fmt.Println("Invalid modified date format. Searching without modified date filter.")
		} else {
			modifiedAfter = parsedDate
		}
	}

	// Call the search function
	results, err := searchFiles(searchDir, fileName, fileType, modifiedAfter)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Display the search results
	fmt.Println("Search Results:")
	for _, file := range results {
		fmt.Println(file)
	}
}

func getInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	_, _ = fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

func searchFiles(searchDir, fileName, fileType string, modifiedAfter time.Time) ([]string, error) {
	var results []string

	//filepath.Walk function in Go is used to traverse a directory tree recursively and perform a function on each visited file or directory
	//For each file or directory visited during the traversal, filepath.Walk will automatically assign the respective path to the path parameter of the callback function
	//filepath.Walk function also provides the info parameter of type os.FileInfo to the callback function for each visited file or directory.
	//The os.FileInfo object contains various metadata and information about the visited file or directory, such as the file size, permissions, modification time, and more.
	err := filepath.Walk(searchDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil // Skip directories
		}

		// Match file name
		if fileName != "" && info.Name() != fileName {
			return nil
		}

		// Match file type
		if fileType != "" && filepath.Ext(info.Name()) != fileType {
			return nil
		}

		// Match modified date
		//info.ModTime() retrieves the modification time of the file or directory being visited
		//!info.ModTime().After(modifiedAfter) is a condition that checks if the modification time of a file is not after the provided modifiedAfter time
		//the program filters files based on whether their modification time is after the specified date
		if !modifiedAfter.IsZero() && !info.ModTime().After(modifiedAfter) {
			return nil
		}

		// Add the matched file to results
		results = append(results, path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}
