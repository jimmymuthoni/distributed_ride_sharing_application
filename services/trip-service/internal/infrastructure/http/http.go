package http

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/jimmymuthoni/distributed_ride_sharing_application/shared/types"
	"github.com/jimmymuthoni/distributed_ride_sharing_application/services/trip-service/internal/domain"
)

type HttpHandler struct {
	Service domain.TripService
}


type previewTripRequest struct {
	UserID		string			 `json:"userID"` 	
	Pickup		types.Coordinate `json:"pickup"`
	Destination	types.Coordinate `json:"destination"`
}

func (s *HttpHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		http.Error(w,"failed to parse JSON data",http.StatusBadRequest)
		return
	}


	ctx := r.Context()

	t, err := s.Service.GetRoute(ctx,&reqBody.Pickup, &reqBody.Destination)
	if err != nil {
		log.Println(err)
	}
	
	writeJson(w,http.StatusOK, t)

}

func writeJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}