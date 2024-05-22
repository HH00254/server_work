package main

import (
	"log"

	"github.com/HH00254/server_work/internal/server"
)

func main() {

	srv := server.NewServer()
	// go func() {
	// 	err := srv.ListenAndServe()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// }()

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
