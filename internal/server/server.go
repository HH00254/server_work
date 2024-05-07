package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func NewServer() *http.Server {

	router := chi.NewRouter()
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "public"))
	fs := http.FileServer(filesDir)

	router.Get("/isAlive", func(w http.ResponseWriter, r *http.Request) {
		responWithJSON(w, 200, struct{}{})
	})

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + "8080",
	}

	return srv
}

func responWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	byteData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)

	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(byteData)
}
