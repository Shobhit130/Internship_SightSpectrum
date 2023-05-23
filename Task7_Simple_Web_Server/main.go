package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Specify the directory to serve files from
	dir := "./static"

	// Create a file server handler
	fileServer := http.FileServer(http.Dir(dir))

	// Create a custom handler to serve the files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check if the requested URL path is the root path ("/")
		if r.URL.Path == "/" {
			// Display custom message
			message := "Type any file name in the URL"
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(message))
			return
		}

		// Get the requested file path
		filePath := filepath.Join(dir, r.URL.Path)

		// Check if the file exists
		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			// File not found, return a 404 error
			//setting the response status code to 404 (Not Found)
			http.NotFound(w, r)
			return
		}

		// serving the requested file
		fileServer.ServeHTTP(w, r)
	})

	// Start the server
	port := 8080
	host := "localhost"
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
