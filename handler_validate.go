package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type cleanedBody struct {
    CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}
	
	cleaned := cleanBody(params.Body)
	respondWithJSON(w, http.StatusOK, cleanedBody{
		CleanedBody: cleaned,
	})
}

func cleanBody(body string) string {
	banned := []string{"kerfuffle", "sharbert", "fornax"}

	words := strings.Fields(body)

	for i, w := range words {
		lower := strings.ToLower(w)
		for _, b := range banned {
			if lower == b {
				words[i] = "****"
				break
			}
		}
	}
	return strings.Join(words, " ")
}

