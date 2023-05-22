package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// URL of the website to scrape
	url := "https://www.bbc.co.uk/news"

	// Send HTTP GET request to the website
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}
	
	// Store the articles in a CSV file
	fileName := "articles.csv"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write column headers to the CSV file
	headers := []string{"Headline","Summary"}
	err = writer.Write(headers)
	if err != nil {
		panic(err)
	}

	// Extract news articles
	articles := [][]string{}
	doc.Find(".gs-c-promo").Each(func(i int, s *goquery.Selection) {
		headline := strings.TrimSpace(s.Find(".gs-c-promo-heading__title").Text())
		summary := strings.TrimSpace(s.Find(".gs-c-promo-summary").Text())
		article := []string{headline, summary}
		articles = append(articles, article)
	})


	for _, article := range articles {
		err := writer.Write(article)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("Data has been successfully stored in %s\n", fileName)
}
