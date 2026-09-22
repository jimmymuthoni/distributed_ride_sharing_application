package main

import (
	"log"
	"net/http"

	"github.com/jimmymuthoni/distributed_ride_sharing_application/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway")

	mux := http.NewServeMux()


	mux.HandleFunc("POST /trip/preview", handleTripPreview)

	server := &http.Server{
		Addr: httpAddr,
		Handler: mux,
	}


	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP seevrr error: %v", err)
	}
}
