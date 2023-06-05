package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/disintegration/imaging"
)

func main() {
	// Prompt the user to enter the input image path
	fmt.Print("Enter the input image path: ")
	var inputImagePath string
	fmt.Scanln(&inputImagePath)

	// Prompt the user to enter the desired width
	fmt.Print("Enter the desired width: ")
	var widthStr string
	fmt.Scanln(&widthStr)
	width, err := strconv.Atoi(widthStr)
	if err != nil {
		log.Fatal("Invalid width:", err)
	}

	// Prompt the user to enter the desired height
	fmt.Print("Enter the desired height: ")
	var heightStr string
	fmt.Scanln(&heightStr)
	height, err := strconv.Atoi(heightStr)
	if err != nil {
		log.Fatal("Invalid height:", err)
	}

	// Specify the output image path
	outputImagePath := "output.jpg"

	// Open the input image file
	inputFile, err := os.Open(inputImagePath)
	if err != nil {
		log.Fatal("Failed to open input image:", err)
	}
	defer inputFile.Close()

	// Decode the input image
	inputImage, err := imaging.Decode(inputFile)
	if err != nil {
		log.Fatal("Failed to decode input image:", err)
	}

	// Resize the image to the desired dimensions
	resizedImage := imaging.Resize(inputImage, width, height, imaging.Lanczos)

	// Save the resized image to the output file
	err = imaging.Save(resizedImage, outputImagePath)
	if err != nil {
		log.Fatal("Failed to save output image:", err)
	}

	fmt.Println("Image resized and saved successfully.")
}
