// This file keeps tracks of in-memory variables/data
// For example: Rooms, Players (Websocket connections), etc...

package handlers

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Player struct {
	Name string          `json:"name"`
	Conn *websocket.Conn `json:"-"`
}

// TODO: Add a maximum player cap
type Room struct {
	Code    string   `json:"code"`
	Players []Player `json:"players"`
}

// Stores all the rooms
type RoomManager struct {
	Rooms map[string]*Room // Stores the room code basically for lookup
	sync.RWMutex
}

var Manager = RoomManager{
	Rooms: make(map[string]*Room),
}

// Method to add a player to a room
func (m *RoomManager) AddPlayerToRoom(roomCode string, player Player) {

	// Maintains ACID compliance (Database Systems Moodule paying off)
	m.Lock()
	defer m.Unlock()

	// Finds room, or creates one if not found
	room, exists := m.Rooms[roomCode]
	if !exists {
		room = &Room{Code: roomCode, Players: []Player{}}
		m.Rooms[roomCode] = room
	}

	room.Players = append(room.Players, player) // Makes player join room
}

// Lists all rooms, returns pointers for efficiency
func (m *RoomManager) ListRooms() []Room {
	m.RLock()
	defer m.RUnlock()

	var rooms []Room
	for _, r := range m.Rooms {
		rooms = append(rooms, *r)
	}

	return rooms
}
