package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/HH00254/server_work/internal/jsonFormat"
	"github.com/ProtonMail/gopenpgp/v2/helper"
)

func ReadyCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server is Healthy")
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

func PublicKeyDecryption(w http.ResponseWriter, r *http.Request) {
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

	privateKey := os.Getenv("PRVKEY")

	decrypted, err := helper.DecryptMessageArmored(privateKey, nil, param.EncryptedKey)
	if err != nil {
		panic(err)
	}

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
