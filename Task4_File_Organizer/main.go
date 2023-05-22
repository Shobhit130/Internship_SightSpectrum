package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// Define the source directory
	sourceDir := "./source"

	// Define the destination directories for each file type
	destinationDirs := map[string]string{
		".jpg":  "./Images",
		".png":  "./Images",
		".gif":  "./Images",
		".txt":  "./Documents",
		".pdf":  "./Documents",
		".docx": "./Documents",
		".xlsx": "./Documents",
		".pptx": "./Documents",
		".md":   "./Documents",
		".go":   "./Code",
		".java": "./Code",
		".cpp":  "./Code",
		".py":   "./Code",
		".html": "./Code",
		".css":  "./Code",
		".js":   "./Code",
		".zip":  "./Archives",
		".rar":  "./Archives",
	}

	// Create destination directories if they don't exist
	for _, dir := range destinationDirs {
		//os.MkdirAll creates the directory even if some parent directories don't exist
		//ModePerm - permission mode for the created directory,grants read, write, and execute permissions to the owner
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
	}

	// Get a list of files in the source directory
	files, err := os.ReadDir(sourceDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	// Iterate over each file and move it to the appropriate destination directory
	for _, file := range files {
		if !file.IsDir() {
			filename := file.Name()
			fileExt := filepath.Ext(filename)

			// Check if the file extension has a corresponding destination directory
			destinationDir, ok := destinationDirs[fileExt]
			if !ok {
				destinationDir = "./Other"
			}

			// construct the source path and destination path for a file based on the given directory and filename.
			sourcePath := filepath.Join(sourceDir, filename)
			destinationPath := filepath.Join(destinationDir, filename)

			// Move the file to the destination directory
			err := moveFile(sourcePath, destinationPath)
			if err != nil {
				fmt.Printf("Error moving file: %v\n", err)
			} else {
				fmt.Printf("Moved file %s to %s\n", filename, destinationDir)
			}
		}
	}
}

// Helper function to move a file from source to destination
func moveFile(sourcePath, destinationPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	err = sourceFile.Close()
	if err != nil {
		return err
	}

	err = os.Remove(sourcePath)
	if err != nil {
		return err
	}

	return nil
}
