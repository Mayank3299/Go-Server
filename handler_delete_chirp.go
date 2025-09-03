package main

import (
	"net/http"

	"github.com/Mayank3299/Go-Server/internal/auth"
	"github.com/google/uuid"
)

func (ac *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {

	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	userId, err := auth.ValidateJWT(tokenString, ac.secret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid token", err)
		return
	}

	chirpId, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse id", err)
		return
	}

	chirp, err := ac.db.GetChirp(r.Context(), chirpId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}

	if userId != chirp.UserID.UUID {
		respondWithError(w, http.StatusForbidden, "Forbidden", err)
		return
	}

	err = ac.db.DeleteChirp(r.Context(), chirp.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't remove chirp", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
