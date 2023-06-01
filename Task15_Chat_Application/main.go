package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type Client struct {
	conn     *websocket.Conn
	send     chan *Message
	chat     *Chat
	username string
}

func (c *Client) readMessages() {
	defer func() {
		c.chat.unregister <- c
		c.conn.Close()
	}()

	for {
		var message Message
		err := c.conn.ReadJSON(&message)
		if err != nil {
			// Ignore the error if it's due to WebSocket connection close
			if websocket.IsCloseError(err, websocket.CloseGoingAway) {
				break
			}
			log.Printf("error reading message: %v", err)
			break
		}

		c.chat.broadcast <- &message
	}
}

func (c *Client) writeMessages() {
	defer func() {
		c.conn.Close()
	}()

	for message := range c.send {
		err := c.conn.WriteJSON(message)
		if err != nil {
			log.Printf("error writing message: %v", err)
			break
		}
	}
}

type Chat struct {
	clients    map[*Client]bool
	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client
	mutex      sync.Mutex
}

func (chat *Chat) run() {
	for {
		select {
		case client := <-chat.register:
			chat.mutex.Lock()
			chat.clients[client] = true
			chat.mutex.Unlock()

		case client := <-chat.unregister:
			chat.mutex.Lock()
			delete(chat.clients, client)
			close(client.send)
			chat.mutex.Unlock()

		case message := <-chat.broadcast:
			chat.mutex.Lock()
			for client := range chat.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(chat.clients, client)
				}
			}
			chat.mutex.Unlock()
		}
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(chat *Chat, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket upgrade error: %v", err)
		return
	}

	client := &Client{
		conn: conn,
		send: make(chan *Message),
		chat: chat,
	}

	chat.register <- client

	go client.readMessages()
	go client.writeMessages()
}

func main() {
	chat := &Chat{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}

	go chat.run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocketHandler(chat, w, r)
	})

	fmt.Println("Chat server is running. Open http://localhost:8000 in your web browser.")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
