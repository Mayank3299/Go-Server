package main

import (
	"fmt"
	"net/http"
)

func (ac *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	resp := fmt.Sprintf("Hits: %d", ac.fileserverHits.Load())
	w.Write([]byte(resp))
}
