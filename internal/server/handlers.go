package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/HH00254/server_work/internal/jsonFormat"
)

// The point of this is to hold onto dependancies
func ReadyCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server is Healthy")
	jsonFormat.RespondWithJSON(w, 200, struct{}{})

}

func (c Controller) GetPublicKey(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"public_key": c.publicKey,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)

}

func (c Controller) PublicKeyDecryption(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		EncryptedKey string `json:"encryptedKey"`
	}
	param := &parameters{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(param)
	if err != nil {
		// TODO
		return
	}

	decrypted, err := c.decryptor.Decrypt(param.EncryptedKey)
	if err != nil {
		log.Println("Decryption Error: ", err)
	}
	fmt.Println(decrypted)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("TODO")
}

func RouteToClientPage(w http.ResponseWriter, r *http.Request) {

	var pathNotFound bool = true
	var publicDir string = "./public"
	var filesDir http.Dir = http.Dir(publicDir)
	var fs http.Handler = http.FileServer(filesDir)

	requestPath := filepath.Join(publicDir, r.URL.Path)

	fmt.Println(requestPath)

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
