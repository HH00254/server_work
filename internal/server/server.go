package server

import (
	"net/http"
	"os"

	handler "github.com/HH00254/server_work/internal/handlers"
	"github.com/HH00254/server_work/internal/sessions"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func NewServer() *http.Server {
	godotenv.Load()
	port := os.Getenv("PORT")

	router := chi.NewRouter()

	router.Get("/isAlive", handler.ReadyCheck)
	router.Get("/session", sessions.MySessionHandler)
	router.Get("/pgkey", handler.GetPublicKey)
	router.Post("/decryption", handler.PublicKeyDecryption)
	router.Get("/*", handler.RouteToClientPage)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	return srv
}
