package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jimmymuthoni/distributed_ride_sharing_application/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request){

	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w,"failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//validation
	if reqBody.UserID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return	
	}

	jsonBody, _ := json.Marshal(reqBody)
	reader := bytes.NewReader(jsonBody)
	
	//call trip service
	resp, err := http.Post("http://trip-service:8083/preview","application/json", reader)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()

	var respBody any
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		http.Error(w,"failed to parse JSON data on trip service", http.StatusBadRequest)
		return
	}
	 

	response := contracts.APIResponse{Data: respBody}
	
	writeJson(w, http.StatusCreated, response)
	

}