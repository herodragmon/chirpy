package main
import (
	"net/http"
	"github.com/google/uuid"
	"sort"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirps", err)
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
	sortString := r.URL.Query().Get("sort")
	responseChirps := []Chirp{}
	for _, dbChirp := range chirps {
		if authorID != uuid.Nil && dbChirp.UserID != authorID {
			continue
		}
		responseChirp := Chirp{
    ID:        dbChirp.ID,
    CreatedAt: dbChirp.CreatedAt,
    UpdatedAt: dbChirp.UpdatedAt,
    Body:      dbChirp.Body,
    UserID:    dbChirp.UserID,
		}
		responseChirps = append(responseChirps, responseChirp)
	}
	if sortString == "desc" {
		sort.Slice(responseChirps, func(i, j int) bool {
			return responseChirps[i].CreatedAt.After(responseChirps[j].CreatedAt)
		})
	} else {
		sort.Slice(responseChirps, func(i, j int) bool {
			return responseChirps[i].CreatedAt.Before(responseChirps[j].CreatedAt)
		})
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
