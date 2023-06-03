package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func getHtml(url string) *http.Response {
	response, err := http.Get(url)
	checkNilError(err)

	if response.StatusCode > 400 {
		fmt.Println("Status Code: ", response.StatusCode)
	}

	return response
}

func writeData(data [][]string) {
	file, err := os.Create("output.csv")
	checkNilError(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header
	header := []string{"Title", "Price", "Link"}
	err = writer.Write(header)
	checkNilError(err)

	// Write the data rows
	for _, row := range data {
		err := writer.Write(row)
		checkNilError(err)
	}
}

func scrapePageData(doc *goquery.Document) {
	var data [][]string

	doc.Find("ul.srp-results>li.s-item").Each(func(i int, s *goquery.Selection) {
		//link of each product
		aLink := s.Find("a.s-item__link")
		url, _ := aLink.Attr("href")

		//title of each product
		titleDiv := s.Find("a.s-item__link>div.s-item__title")
		title := strings.TrimSpace(titleDiv.Text())

		//price of each product
		priceSpan := strings.TrimSpace(s.Find("span.s-item__price").Text())

		data = append(data, []string{title, priceSpan, url})
	})
	writeData(data)
}

func main() {
	var url string
	fmt.Print("Enter the URL: ")
	fmt.Scan(&url)

	response := getHtml(url)
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	checkNilError(err)

	scrapePageData(doc)

	fmt.Println("Data saved to output.csv")
}
