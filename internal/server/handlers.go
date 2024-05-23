package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// The point of this is to hold onto dependancies

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
