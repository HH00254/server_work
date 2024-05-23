package server

import (
	"net/http"

	handler "github.com/HH00254/server_work/internal/handlers"
	"github.com/HH00254/server_work/internal/sessions"
	"github.com/HH00254/server_work/util"
	"github.com/go-chi/chi/v5"
)

type Controller struct {
	publicKey string
	decryptor util.PgpDecryptor
}

func NewController(publicKey string, decryptor util.PgpDecryptor) Controller {
	return Controller{
		publicKey: publicKey,
		decryptor: decryptor,
	}
}

func NewServer(port, publicKey, privateKey string) *http.Server {

	controller := NewController(publicKey, util.NewPgpDecrypter(privateKey))
	router := chi.NewRouter()

	// pattern dependency injection in main.go normally
	router.Get("/isAlive", handler.ReadyCheck)
	router.Get("/session", sessions.MySessionHandler)
	router.Get("/pgkey", handler.GetPublicKey)
	router.Post("/decryption", controller.PublicKeyDecryption)
	router.Get("/*", handler.RouteToClientPage)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	return srv
}
