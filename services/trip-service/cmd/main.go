package main

import (
	"context"
	"log"
	"time"

	"github.com/jimmymuthoni/distributed_ride_sharing_application/services/trip-service/internal/domain"
	repository "github.com/jimmymuthoni/distributed_ride_sharing_application/services/trip-service/internal/infrastructure"
	"github.com/jimmymuthoni/distributed_ride_sharing_application/services/trip-service/internal/service"
)

func main(){

	ctx := context.Background()
	inmem := repository.NewInmemRepository()
	svc := service.NewService(inmem)

	fare := &domain.RideFareModel{
		UserID: "42",
	}
	t, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}
    log.Println(t)

	//keep the program runnning for now
	for {
		time.Sleep(time.Second)
	}

}

