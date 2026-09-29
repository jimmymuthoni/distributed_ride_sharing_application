package main

import (
	"log"
	"net/http"
	repository "github.com/jimmymuthoni/distributed_ride_sharing_application/services/trip-service/internal/infrastructure"
	"github.com/jimmymuthoni/distributed_ride_sharing_application/services/trip-service/internal/service"
	h "github.com/jimmymuthoni/distributed_ride_sharing_application/services/trip-service/internal/infrastructure/http"
)

 

func main(){
	log.Println("Starting Trip service server....")
	inmem := repository.NewInmemRepository()
	svc := service.NewService(inmem)

	mux := http.NewServeMux()
	httphandler := h.HttpHandler{Service: svc}

	mux.HandleFunc("POST /preview", httphandler.HandleTripPreview)

	server := &http.Server{
		Addr: ":8083",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error:%v", err)
	}


}


