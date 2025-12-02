package main
import (
	"net/http"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirps", err)
		return
	}
	responseChirps := []Chirp{}
	for _, dbChirp := range chirps {
		responseChirp := Chirp{
    ID:        dbChirp.ID,
    CreatedAt: dbChirp.CreatedAt,
    UpdatedAt: dbChirp.UpdatedAt,
    Body:      dbChirp.Body,
    UserID:    dbChirp.UserID,
		}
		responseChirps = append(responseChirps, responseChirp)
	}
	respondWithJSON(w, http.StatusOK, responseChirps)
}

func (cfg *apiConfig) handlerChirpsGetByID(w http.ResponseWriter, r *http.Request) {
	chirpIDStr := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid id", err)
		return
	}
	chirp, err := cfg.db.GetChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp", err)
		return
	}
	respondWithJSON(w, http.StatusOK, Chirp{
    ID:        chirp.ID,
    CreatedAt: chirp.CreatedAt,
    UpdatedAt: chirp.UpdatedAt,
    Body:      chirp.Body,
    UserID:    chirp.UserID,
	})
}
