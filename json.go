package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, statusCode int, message string, err error) {
	if err != nil {
		log.Println(err)
	}
	if statusCode > 499 {
		log.Printf("Responding with 5XX error: %s", message)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, statusCode, errorResponse{
		Error: message + err.Error(),
	})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(dat)
}
