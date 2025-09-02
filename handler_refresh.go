package main

import (
	"net/http"
	"time"

	"github.com/Mayank3299/Go-Server/internal/auth"
)

func (ac *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}

	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	refreshToken, err := ac.db.GetRefreshToken(r.Context(), tokenString)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid token value", err)
		return
	}

	expiresAt := refreshToken.ExpiresAt
	if expiresAt.Before(time.Now().UTC()) {
		respondWithError(w, http.StatusUnauthorized, "token expired", err)
		return
	}

	revokedAt := refreshToken.RevokedAt
	if revokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "token revoked", err)
		return
	}

	userId := refreshToken.UserID.UUID
	accessToken, err := auth.MakeJWT(userId, ac.secret)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't make JWT", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		Token: accessToken,
	})
}
