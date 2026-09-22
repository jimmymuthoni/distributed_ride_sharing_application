package main

import (
	"github.com/jimmymuthoni/distributed_ride_sharing_application/shared/types"
	
)


type previewTripRequest struct {
	UserID		string			 `json:"userID"` 	
	Pickup		types.Coordinate `json:"pickup"`
	Destination	types.Coordinate `json:"destination"`
}