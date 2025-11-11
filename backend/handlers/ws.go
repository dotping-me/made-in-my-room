package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Just a simple function that prints back any messages it receives for now
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error on Websocket connection:", err)
		return
	}

	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Could not read Websocket message:", err)
			break // Most likely a disconnection
		}

		receivedMessage := string(msg)
		fmt.Println("Received:", receivedMessage) // Converts bytes
		conn.WriteMessage(websocket.TextMessage, []byte(receivedMessage))
	}
}
