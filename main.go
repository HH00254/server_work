package main

import (
	"log"
	"os"

	"github.com/HH00254/server_work/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	srv := server.NewServer(
		os.Getenv("PORT"),
		os.Getenv("PUBKEY"),
		os.Getenv("PRVKEY"),
	)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
