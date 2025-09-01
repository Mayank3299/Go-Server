package main

import (
	"net/http"
)

func (ac *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if ac.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Reset is only allowed in dev environment."))
		return
	}
	ac.fileserverHits.Store(0)
	err := ac.db.DeleteUsers(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to reset the database: " + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0 and database reset to initial state."))
}
