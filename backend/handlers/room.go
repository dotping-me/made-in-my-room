package handlers

import (
	"encoding/json"
	"net/http"
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
