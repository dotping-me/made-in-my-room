package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dotping-me/made-in-my-room/utils"
)

// Checks if a room with specified code exists
func DoesRoomExist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	roomCode := r.URL.Query().Get("room")
	if roomCode == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Room Code missing!"})
		return
	}

	_, exists := Manager.Rooms[roomCode]
	if !exists {
		json.NewEncoder(w).Encode(map[string]bool{"exists": false})
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"exists": true})
}

// Creates a room and returns room code
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// NOTE: Do I add a loop to retry until a unique code is CERTAINLY generated?
	// Or are the codes unique enough on their own?

	code := utils.GenerateRandomCode(4)
	_, exists := Manager.Rooms[code]

	if code == "" || exists {
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to Generate Code!"})
		return
	}

	// Creates room
	Manager.Rooms[code] = &Room{Code: code, Players: []Player{}, Cap: 8}
	fmt.Println("Created Room", code)

	json.NewEncoder(w).Encode(map[string]string{"code": code})
}
