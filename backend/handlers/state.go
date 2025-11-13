package handlers

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

// NOTE: Using in-memory storgare for all this,
// 		 it should be fine if the player count is small but then we might
// 		 might have to switch to using a database solution

type Player struct {
	Name string          `json:"name"`
	Conn *websocket.Conn `json:"-"`
}

// TODO: Add a maximum player cap
type Room struct {
	Code    string   `json:"code"`
	Cap     uint8    `json:"max_size"`
	Players []Player `json:"players"`
	Host    Player   `json:"host"`
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

	} else {
		if room.Cap == 8 {
			return // Check if room is full
		}
	}

	room.Players = append(room.Players, player) // Makes player join room
	room.Cap++

	// Set host
	if len(room.Players) == 1 {
		room.Host = player
	}

	fmt.Println("Added Player [", player.Name, "] to Room", room.Code)
}

func (m *RoomManager) RemovePlayerFromRoom(roomCode string, player Player) {
	m.Lock()
	defer m.Unlock()

	room, exists := m.Rooms[roomCode]
	if !exists {
		return
	}

	remainingPlayers := []Player{}
	for _, p := range room.Players {
		if p.Conn != player.Conn {
			remainingPlayers = append(remainingPlayers, p)
		}
	}

	// Delete room if there are no players in it
	if len(remainingPlayers) == 0 {
		delete(Manager.Rooms, roomCode)
		return
	}

	room.Players = remainingPlayers
	room.Cap--

	// Updates host
	if len(room.Players) == 1 {
		room.Host = player
	}

	fmt.Println("Removed Player [", player.Name, "] to Room", room.Code)
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
