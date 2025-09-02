package main

import (
	"net/http"

	"github.com/Mayank3299/Go-Server/internal/auth"
)

func (ac *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	_, err = ac.db.GetRefreshToken(r.Context(), tokenString)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid token value", err)
		return
	}

	err = ac.db.RevokeToken(r.Context(), tokenString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't revoke token", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
