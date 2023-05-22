package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Get the API key 
	apiKey := "8009b4b6a80843139512dc5a84908270"

	// Get input from the user
	var amount float64
	var sourceCurrency, targetCurrency string
	fmt.Println("Enter the amount:")
	fmt.Scan(&amount)
	fmt.Println("Enter the source currency:")
	fmt.Scan(&sourceCurrency)
	fmt.Println("Enter the target currency:")
	fmt.Scan(&targetCurrency)

	// Build the API URL
	apiURL := fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s&symbols=%s,%s", apiKey, sourceCurrency, targetCurrency)

	// Make the API request
	resp, err := http.Get(apiURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse the response JSON to myData
	var myData map[string]interface{}
	err = json.Unmarshal(body, &myData)
	if err != nil {
		panic(err)
	}

	//Get the rates field of the myData containing the requested currencies' exchange rates
	ratesMap := myData["rates"].(map[string]interface{})

	// Calculate the conversion
	targetRate := ratesMap[targetCurrency].(float64)
    sourceRate := ratesMap[sourceCurrency].(float64)

	convertedAmount := amount * (targetRate / sourceRate)

	// Output the result
	fmt.Printf("%f %s is equivalent to %f %s\n", amount, sourceCurrency, convertedAmount, targetCurrency)
}
