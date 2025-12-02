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
	jwtSecret string
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
}

func main() {
	godotenv.Load()
	platform := os.Getenv("PLATFORM")
	secretKey := os.Getenv("SECRET_KEY")
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
		jwtSecret: secretKey,
	}

	fsHandler := http.FileServer(http.Dir(filepathRoot))
	wrapped := (&cfg).middlewareMetricsInc(fsHandler)

	mux.Handle("/app/", http.StripPrefix("/app/", wrapped))
	mux.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", cfg.handlerReset)
	mux.HandleFunc("POST /api/users", cfg.handlerUsersCreate)
	mux.HandleFunc("POST /api/chirps", cfg.handlerChirpsCreate)
	mux.HandleFunc("GET /api/chirps", cfg.handlerChirpsGet)
	mux.HandleFunc("GET /api/chirps/{chirpID}", cfg.handlerChirpsGetByID)
	mux.HandleFunc("POST /api/login", cfg.handlerLogin)
	mux.HandleFunc("POST /api/refresh", cfg.handlerRefresh)
	mux.HandleFunc("POST /api/revoke", cfg.handlerRevoke)


	server := &http.Server{
		Handler: mux,
		Addr: ":" + port,
	}
	
	log.Printf("Server starting on port %s", port)
	log.Fatal(server.ListenAndServe())
}
