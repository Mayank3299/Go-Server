package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (ac *apiConfig) handlerPolka(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		} `json:"data"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	switch params.Event {
	case "user.upgraded":
		_, err := ac.db.UpgradeUserToRed(r.Context(), params.Data.UserID)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Couldn't upgrade user", err)
			return
		}

		respondWithJSON(w, http.StatusNoContent, nil)
		return
	default:
		respondWithJSON(w, http.StatusNoContent, nil)
		return
	}
}
