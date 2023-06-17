package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/dhowden/tag"
	"github.com/unidoc/unipdf/v3/model"
)

type FileInfo struct {
	FilePath        string
	FileSize        int64
	CreationTime    time.Time
	FileType        string
	ImageWidth      int
	ImageHeight     int
	AudioArtist     string
	AudioAlbum      string
	PDFCreator      string
	PDFAuthor       string
	PDFModifiedDate time.Time
	// Add more fields for other metadata you want to extract
}

func main() {
	var filePath string
	fmt.Print("Enter the file path: ")
	fmt.Scanln(&filePath)

	fileInfo, err := GetFileInfo(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Path:", fileInfo.FilePath)
	fmt.Println("File Size:", fileInfo.FileSize, "bytes")
	fmt.Println("Creation Time:", fileInfo.CreationTime)
	fmt.Println("File Type:", fileInfo.FileType)

	// Print image-specific metadata
	if fileInfo.FileType == "image" {
		fmt.Println("Image Width:", fileInfo.ImageWidth)
		fmt.Println("Image Height:", fileInfo.ImageHeight)
	}

	// Print audio-specific metadata
	if fileInfo.FileType == "audio" {
		fmt.Println("Artist:", fileInfo.AudioArtist)
		fmt.Println("Album:", fileInfo.AudioAlbum)
	}

	// Print PDF-specific metadata
	if fileInfo.FileType == "pdf" {
		fmt.Println("Creator:", fileInfo.PDFCreator)
		fmt.Println("Author:", fileInfo.PDFAuthor)
		fmt.Println("Modified Date:", fileInfo.PDFModifiedDate)
	}

	// Add more print statements for additional metadata fields
}

func GetFileInfo(filePath string) (FileInfo, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return FileInfo{}, err
	}

	//size of the file in bytes
	fileSize := info.Size()
	//the last modification time of the file
	creationTime := info.ModTime()

	fileType, err := getFileType(filePath)
	if err != nil {
		return FileInfo{}, err
	}

	var imageWidth, imageHeight int
	if fileType == "image" {
		imageWidth, imageHeight, err = getImageDimensions(filePath)
		if err != nil {
			return FileInfo{}, err
		}
	}

	var audioArtist, audioAlbum string
	if fileType == "audio" {
		audioArtist, audioAlbum, err = getAudioMetadata(filePath)
		if err != nil {
			return FileInfo{}, err
		}
	}

	var pdfCreator, pdfAuthor string
	var pdfModifiedDate time.Time
	if fileType == "pdf" {
		pdfCreator, pdfAuthor, pdfModifiedDate, err = getPDFMetadata(filePath)
		if err != nil {
			return FileInfo{}, err
		}
	}

	fileInfo := FileInfo{
		FilePath:        filePath,
		FileSize:        fileSize,
		CreationTime:    creationTime,
		FileType:        fileType,
		ImageWidth:      imageWidth,
		ImageHeight:     imageHeight,
		AudioArtist:     audioArtist,
		AudioAlbum:      audioAlbum,
		PDFCreator:      pdfCreator,
		PDFAuthor:       pdfAuthor,
		PDFModifiedDate: pdfModifiedDate,
		// Assign more values to other metadata fields as needed
	}

	return fileInfo, nil
}

func getFileType(filePath string) (string, error) {
	ext := filepath.Ext(filePath)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
		return "image", nil
	case ".mp3", ".flac", ".m4a", ".wav", ".wma":
		return "audio", nil
	case ".pdf":
		return "pdf", nil
	default:
		return "", fmt.Errorf("unknown file type for extension %s", ext)
	}
}

func getImageDimensions(filePath string) (int, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	// decode the image configuration from the given file
	//img contains the decoded image configuration, including information such as width, height, and color model
	//_ is the blank identifier used to discard the color model
	img, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println(filePath)
		return 0, 0, err
	}

	width := img.Width
	height := img.Height

	return width, height, nil
}

func getAudioMetadata(filePath string) (string, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	// extract audio metadata from the specified file
	m, err := tag.ReadFrom(file)
	if err != nil {
		return "", "", err
	}

	artist := m.Artist()
	album := m.Album()

	return artist, album, nil
}

func getPDFMetadata(filePath string) (string, string, time.Time, error) {
	// Open the PDF file.
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening PDF file: %v", err)
		return "", "", time.Time{}, err
	}
	defer f.Close()

	// Create a new PDF reader.
	reader, err := model.NewPdfReader(f)
	if err != nil {
		fmt.Printf("Error reading PDF: %v", err)
		return "", "", time.Time{}, err
	}

	// Get the document info dictionary.
	infoDict, err := reader.GetPdfInfo()
	if err != nil {
		fmt.Printf("Error retrieving document info: %v", err)
		return "", "", time.Time{}, err
	}

	// Extract metadata from the document info dictionary.
	//Creator field represents the software or tool that was used to create the PDF
	creator := infoDict.Creator
	author := infoDict.Author
	modified := infoDict.ModifiedDate

	return creator.Decoded(), author.Decoded(), modified.ToGoTime(), nil
}
