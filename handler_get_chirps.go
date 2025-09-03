package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (ac *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := ac.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Coudln't fetch chirps", err)
		return
	}

	authorID := uuid.Nil
	authorIDString := r.URL.Query().Get("author_id")
	if authorIDString != "" {
		authorID, err = uuid.Parse(authorIDString)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author ID", err)
			return
		}
	}

	chirps := []Chirp{}
	for _, c := range dbChirps {
		if authorID != uuid.Nil && c.UserID.UUID != authorID {
			continue
		}
		chirps = append(chirps, Chirp{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			Body:      c.Body,
			UserID:    c.UserID.UUID,
		})
	}

	respondWithJSON(w, http.StatusOK, chirps)
}
