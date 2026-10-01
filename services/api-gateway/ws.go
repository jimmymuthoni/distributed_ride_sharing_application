package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jimmymuthoni/distributed_ride_sharing_application/shared/contracts"
	"github.com/jimmymuthoni/distributed_ride_sharing_application/shared/util"
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
		log.Printf("Error reading message: %v", err)
		break 
	}

	log.Printf("Received message: %s", message)

	}

}


func handleDriversWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Websocket upgrade failed: %v", err)
		return
	}
	defer conn.Close()
	
	userId := r.URL.Query().Get("userID")
	if userId == "" {
		log.Printf("UserId is not provided")
		return
	}

	packageSlug := r.URL.Query().Get("packageSlug")
	if packageSlug == "" {
		log.Printf("Packageslug is not provided")
		return
	}

	type Driver struct{
		Id				string `json:"id"`
		Name			string `json:"name"`
		ProfilePicture	string `json:"profilePicture"`
		CarPlate		string `json:"carPlate"`
		PackageSlug		string `json:"packageslug"`
	}

	msg := contracts.WSMessage{
		Type: "driver.cmd.register",
		Data: Driver {
			Id: userId,
			Name: "Jim",
			ProfilePicture: util.GetRandomAvatar(1),
			CarPlate: "abc123",
			PackageSlug: packageSlug,
		},
	}

	if err := conn.WriteJSON(msg); err != nil {
		log.Printf("Error seding message: %v", err)
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
		log.Printf("Error reading message: %v", err)
		break 
	}

	log.Printf("Received message: %s", message)

	}

}