package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var (
	urls    map[string]string
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	baseURL = "http://localhost:8080/"
)

func main() {
	urls = make(map[string]string)
	http.HandleFunc("/", redirectHandler)	
	http.HandleFunc("/shorten", shortenHandler)
	http.ListenAndServe(":8080", nil)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]
	longURL, exists := urls[shortURL]
	if exists {
		http.Redirect(w, r, longURL, http.StatusFound)
	} else {
		fmt.Fprintf(w, "URL not found")
	}
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	longURL := r.FormValue("url")
	shortURL := generateShortURL()
	urls[shortURL] = longURL
	shortenedURL := baseURL + shortURL
	fmt.Fprintf(w, "Short URL: %s", shortenedURL)
}

func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < 8; i++ {
		sb.WriteRune(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}