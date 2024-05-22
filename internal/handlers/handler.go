package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/HH00254/server_work/internal/jsonFormat"
)

func ReadyCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	jsonFormat.RespondWithJSON(w, 200, struct{}{})

}

func GetPublicKey(w http.ResponseWriter, r *http.Request) {
	publicKey := os.Getenv("PUBKEY")

	data := map[string]string{
		"public_key": publicKey,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)

}

func RouteToClientPage(w http.ResponseWriter, r *http.Request) {

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

}
