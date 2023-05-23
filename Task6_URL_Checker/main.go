package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkURL(url string) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		resp = &http.Response{
			StatusCode: http.StatusNotFound,
		}
	}

	fmt.Printf("URL: %s\n", url)
	fmt.Printf("Status code: %d\n", resp.StatusCode)

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		fmt.Println("Accessibility: Accessible")
	} else {
		fmt.Println("Accessibility: Not Accessible")
	}

	fmt.Println()
}

func main() {
	urls := []string{
		"https://www.example.com",
		"https://www.google.com",
		"https://www.invalidurl12345.com",
	}

	for _, url := range urls {
		checkURL(url)
	}
}
