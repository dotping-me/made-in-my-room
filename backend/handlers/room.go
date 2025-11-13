package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dotping-me/made-in-my-room/utils"
)

// NOTE: Using in-memory storgare for all this,
// 		 it should be fine if the player count is small but then we might
// 		 might have to switch to using a database solution

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

	// TODO: Check if room is full too

	json.NewEncoder(w).Encode(map[string]bool{"exists": true})
}

// Creates a room and returns room code
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: Do I add a loop to retry until a unique code is CERTAINLY generated?
	// Or are the codes unique enough on their own?

	code := utils.GenerateRandomCode(4)
	_, exists := Manager.Rooms[code]

	if code == "" || exists {
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to Generate Code!"})
		return
	}

	// Creates room
	Manager.Rooms[code] = &Room{Code: code, Players: []Player{}}
	json.NewEncoder(w).Encode(map[string]string{"code": code})
}

func RoomHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Hardcoded rooms for now
	if len(Manager.Rooms) == 0 {
		Manager.Rooms["1"] = &Room{Code: "1"}
		Manager.Rooms["2"] = &Room{Code: "2"}
	}

	rooms := Manager.ListRooms()
	json.NewEncoder(w).Encode(rooms)
}
