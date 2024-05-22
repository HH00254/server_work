package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/HH00254/server_work/internal/sessions"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func NewServer() *http.Server {
	godotenv.Load()

	router := chi.NewRouter()
	// workDir, _ := os.Getwd()
	// filesDir := http.Dir(filepath.Join(workDir, "public"))
	// look into a way to disable the access to all the files within the root of the public directory for public

	router.Get("/isAlive", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("here")
		responWithJSON(w, 200, struct{}{})
	})

	router.Get("/session", func(w http.ResponseWriter, r *http.Request) {
		sessions.MySessionHandler(w, r)
	})

	router.Get("/pgkey", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Here 34")

		publicKey := os.Getenv("PUBKEY")
		fmt.Print(publicKey)

		data := map[string]string{
			"public_key": publicKey,
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(data)

	})

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {

		var pathNotFound bool = true
		var publicDir string = "./public"
		var filesDir http.Dir = http.Dir(publicDir)
		var fs http.Handler = http.FileServer(filesDir)

		requestPath := filepath.Join(publicDir, r.URL.Path)

		stat, err := os.Stat(requestPath)
		if err == nil && stat.IsDir() {
			pathNotFound = false

			indexPath := filepath.Join(publicDir, "html/index.html")
			http.ServeFile(w, r, indexPath)

		} else if err == nil && pathNotFound {
			pathNotFound = false

			fs.ServeHTTP(w, r)

		} else {

			http.NotFound(w, r)
		}

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
