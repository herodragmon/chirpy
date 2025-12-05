package main

import (
	"github.com/herodragmon/chirpy/internal/auth"
	"encoding/json"
	"net/http"
	"github.com/herodragmon/chirpy/internal/database"

)

func (cfg *apiConfig) handlerUsersUpdate(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	params := request{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find token", err)
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized token", err)
		return
	}
	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
    respondWithError(w, http.StatusInternalServerError, "Couldn't hash password", err)
    return
	}
	user, err := cfg.db.UpdateUser(r.Context(), database.UpdateUserParams{
    Email:          params.Email,
    HashedPassword: hashedPassword,
		ID:             userID,
	})
	if err != nil {
    respondWithError(w, http.StatusInternalServerError, "Couldn't update user", err)
    return
	}
	respondWithJSON(w, http.StatusOK, User{
    ID:          user.ID,
    CreatedAt:   user.CreatedAt,
    UpdatedAt:   user.UpdatedAt,
    Email:       user.Email,
		IsChirpyRed: user.IsChirpyRed.Bool,
	})
}
