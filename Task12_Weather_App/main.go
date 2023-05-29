package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Get the API key
	apiKey := "bc8edf2f2d9076284fee7495b2a92492"

	// Get input from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the city name or ZIP code: ")
	location, _ := reader.ReadString('\n')
	location = strings.TrimSpace(location)

	// Build the API URL
	apiURL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", location, apiKey)

	// Make the API request
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Failed to fetch weather data:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	// Check if the API request was successful
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to fetch weather data. Response:", resp.Status)
		return
	}

	// Parse the response JSON
	var weatherData map[string]interface{}
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println("Failed to parse weather data:", err)
		return
	}

	// Extract the necessary information from the response
	weather, ok := weatherData["weather"].([]interface{})
	if !ok || len(weather) == 0 {
		fmt.Println("Failed to fetch weather data. Invalid response.")
		return
	}

	description := weather[0].(map[string]interface{})["description"].(string)

	main, ok := weatherData["main"].(map[string]interface{})
	if !ok {
		fmt.Println("Failed to fetch weather data. Invalid response.")
		return
	}

	name := weatherData["name"].(string)

	temperature := main["temp"].(float64)
	humidity := main["humidity"].(float64)

	// Output the weather information
	fmt.Println("\nCurrent Weather Conditions:")
	fmt.Printf("Location: %s\n", name)
	fmt.Printf("Description: %s\n", description)
	fmt.Printf("Temperature: %.2f°C\n", temperature-273.15)
	fmt.Printf("Humidity: %.2f%%\n", humidity)
}
