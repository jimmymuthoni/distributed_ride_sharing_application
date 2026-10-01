package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleRidersWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Websocket upgrade failed: %v", err)
		return
	}
	defer conn.Close()
	
	userID := r.URL.Query().Get("userID")
	if userID == "" {
		log.Printf("UserId is not provided")
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
		log.Printf("error reading message: %v", err)
		break 
	}

	log.Printf("received message: %s", message)

	}

}


func handleDriversWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Websocket upgrade failed: %v", err)
		return
	}
	defer conn.Close()
	
	userID := r.URL.Query().Get("userID")
		if userID == "" {
			log.Printf("UserId is not provided")
			return
		}

	}