package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Mayank3299/Go-Server/internal/auth"
	"github.com/Mayank3299/Go-Server/internal/database"
	"github.com/google/uuid"
)

func (ac *apiConfig) handlerLoginUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		User
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	user, err := ac.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	accessToken, err := auth.MakeJWT(user.ID, ac.secret)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't make JWT", err)
		return
	}

	refreshTokenString, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't make refresh token string", err)
		return
	}
	expiresAt := time.Now().UTC().Add(60 * 24 * time.Hour)

	queryParams := database.CreateRefreshTokenParams{
		Token:     refreshTokenString,
		UserID:    uuid.NullUUID{UUID: user.ID, Valid: true},
		ExpiresAt: expiresAt,
	}
	refreshToken, err := ac.db.CreateRefreshToken(r.Context(), queryParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't make refresh token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:          user.ID,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
			Email:       user.Email,
			IsChirpyRed: user.IsChirpyRed,
		},
		Token:        accessToken,
		RefreshToken: refreshToken.Token})
}
