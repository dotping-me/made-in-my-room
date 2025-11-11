package handlers

import (
	"encoding/json"
	"net/http"
)

type Lobby struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func LobbyHandler(w http.ResponseWriter, r *http.Request) {
	lobbies := []Lobby{
		{ID: "1", Name: "Room 1"},
		{ID: "2", Name: "Room 2"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lobbies)
}
