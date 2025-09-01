package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (ac *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}

	type response struct {
		User
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	userEmail := params.Email
	user, err := ac.db.CreateUser(r.Context(), userEmail)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "User not created", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, response{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     userEmail,
		},
	})
}
