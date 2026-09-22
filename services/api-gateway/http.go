package main

import (
	"encoding/json"
	"net/http"

	"github.com/jimmymuthoni/distributed_ride_sharing_application/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request){

	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w,"failed to pare JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//validation
	if reqBody.UserID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return	
	}
	
	//call trip service
	response := contracts.APIResponse{Data: "ok"}
	
	writeJson(w, http.StatusCreated, response)
	

}