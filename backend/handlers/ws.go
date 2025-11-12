package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Sends a Websocket message to all players in the specified room
func broadcast(roomCode string, msgType string, data interface{}) {
	Manager.RLock()

	// Tries to find room
	room, exists := Manager.Rooms[roomCode]
	Manager.RUnlock()

	if !exists {
		return
	}

	// Sends JSON data to all players in that room
	msg, _ := json.Marshal(map[string]interface{}{
		"type": msgType,
		"data": data,
	})

	for _, player := range room.Players {
		if player.Conn != nil {
			player.Conn.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

// Function that handles room comms (Uses GET params just for ease
// of data transfer)
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error on Websocket connection:", err)
		return
	}

	defer conn.Close()

	roomCode := r.URL.Query().Get("room")
	playerName := r.URL.Query().Get("name")
	if roomCode == "" || playerName == "" {
		fmt.Println("Room Code and Player Name are not provided!")
		return
	}

	// Adds player to room and notifies everyone in that room
	player := Player{Name: playerName, Conn: conn}
	Manager.AddPlayerToRoom(roomCode, player)
	broadcast(roomCode, "players", Manager.Rooms[roomCode].Players)

	// Listens and handles all other messages received until disconnection
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Player disconnected:", player.Name)
			break // Most likely a disconnection
		}

		// Received Websocket message (typically JSON)
		var data map[string]interface{}
		if err := json.Unmarshal(msg, &data); err != nil {
			fmt.Println("Invalid JSON from", player.Name)
			continue
		}

		fmt.Printf("[%s] %v\n", player.Name, data)

		// Just echo back any message for now
		if data["type"] == "ping" {
			conn.WriteMessage(websocket.TextMessage, []byte(`{"type":"pong"}`))
		}

		// Remove player from room on disconnect
		Manager.Lock()
		room := Manager.Rooms[roomCode]

		remainingPlayers := []Player{}
		for _, p := range room.Players {
			if p.Conn != conn {
				remainingPlayers = append(remainingPlayers, p)
			}
		}

		room.Players = remainingPlayers
		Manager.Unlock()

		broadcast(roomCode, "players", room.Players)
	}
}
