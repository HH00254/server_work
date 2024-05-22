package handler

import (
	"fmt"
	"net/http"
)

func Ishealthy(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	json.respondWithJSON(w, 200, struct{}{})

}
