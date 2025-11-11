package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/dotping-me/made-in-my-room/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Registering routes
	mux.HandleFunc("/api/lobby", handlers.LobbyHandler)
	mux.HandleFunc("/ws", handlers.WebsocketHandler)

	// Allows Cross-Origin comms
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	}).Handler(mux)

	fmt.Println("mux listening on [http://localhost:8080]")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
