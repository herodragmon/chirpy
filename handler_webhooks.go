package main 

import (
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/herodragmon/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerWebHooks(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Event string `json:"event"`
		Data struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find api", nil)
		return
	}
	if apiKey != cfg.apiPolkaKey {
		respondWithError(w, http.StatusUnauthorized, "Request is not authorized from others", err)
		return
	}
	decoder := json.NewDecoder(r.Body)
	params := request{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	userID, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid id", err)
		return
	}
	_, err = cfg.db.UpgradeUserToChirpyRed(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't find the user", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
