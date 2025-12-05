package main

import (
	"github.com/herodragmon/chirpy/internal/auth"
	"encoding/json"
	"net/http"
	"time"
	"github.com/herodragmon/chirpy/internal/database"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password         string `json:"password"`
		Email            string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	match, err := auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil || !match {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	expirationTime := time.Hour
	
	accessToken, err := auth.MakeJWT(user.ID, cfg.jwtSecret, expirationTime)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't make jwt", err)
		return
	}
	type response struct {
    User
    Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
	
	refreshToken := auth.MakeRefreshToken()

	expiresAt := time.Now().Add(60 * 24 * time.Hour)

	_, err = cfg.db.CreateRefreshTokens(r.Context(), database.CreateRefreshTokensParams{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: expiresAt,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create refresh token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:          user.ID,
			Email:       user.Email,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
			IsChirpyRed: user.IsChirpyRed.Bool,
		},
		Token:        accessToken,
		RefreshToken: refreshToken,
	})
}
