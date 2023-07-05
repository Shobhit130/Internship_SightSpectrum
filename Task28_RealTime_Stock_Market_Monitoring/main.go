package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	apiKey     = "6Q8ICPRMFN815DXS"
	socketPath = "/ws"
	port       = ":8080"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func main() {
	http.HandleFunc(socketPath, handleWebSocket)
	http.HandleFunc("/", handleIndex)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Read the stock symbol from the client
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	stockSymbol := string(message)

	// Fetch real-time stock data
	stockData, err := fetchStockData(stockSymbol)
	if err != nil {
		log.Println(err)
		return
	}

	// Send initial stock data to the client
	err = conn.WriteJSON(stockData)
	if err != nil {
		log.Println(err)
		return
	}

	// Continuously fetch and send updated stock data
	for {
		updatedStockData, err := fetchStockData(stockSymbol)
		if err != nil {
			log.Println(err)
			return
		}

		err = conn.WriteJSON(updatedStockData)
		if err != nil {
			log.Println(err)
			return
		}

		// Delay for 1 minute (60 seconds)
		time.Sleep(60 * time.Second)
	}
}

func fetchStockData(stockSymbol string) (map[string]interface{}, error) {
	// Fetch stock data using the Alpha Vantage API
	// Replace the API call below with your own implementation
	// You need to sign up for an API key at https://www.alphavantage.co/
	// and replace `YOUR_ALPHA_VANTAGE_API_KEY` with your actual API key
	// in the `apiKey` constant at the beginning of the code.
	apiURL := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", stockSymbol, apiKey)
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Serve your HTML/JS/CSS files for the client-side application here
	http.ServeFile(w, r, "index.html")
}
