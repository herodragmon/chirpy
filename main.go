package main

import (
	"log"
	"net/http"
	"sync/atomic"
	_ "github.com/lib/pq"
	"os"
	"database/sql"
	"time"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/herodragmon/chirpy/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db  *database.Queries
	platform string
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func main() {
	godotenv.Load()
	platform := os.Getenv("PLATFORM")
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}
	dbQueries := database.New(db)

	const filepathRoot = "."
	const port = "8080"
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	
	cfg:= apiConfig{
		fileserverHits: atomic.Int32{},
		db: dbQueries,
		platform: platform,
	}

	fsHandler := http.FileServer(http.Dir(filepathRoot))

	wrapped := (&cfg).middlewareMetricsInc(fsHandler)

	mux.Handle("/app/", http.StripPrefix("/app/", wrapped))

	mux.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)

	mux.HandleFunc("POST /admin/reset", cfg.handlerReset)

	mux.HandleFunc("POST /api/validate_chirp", handlerChirpsValidate)

	mux.HandleFunc("POST /api/users", cfg.handlerUsersCreate)


	server := &http.Server{
		Handler: mux,
		Addr: ":" + port,
	}
	
	log.Printf("Server starting on port %s", port)
	log.Fatal(server.ListenAndServe())
}
