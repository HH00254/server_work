package sessions

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

func createSessionStore() *sessions.CookieStore {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sessionKey := os.Getenv("SESSION_KEY")
	var store = sessions.NewCookieStore([]byte(sessionKey))

	return store
}

func MySessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionStore := createSessionStore()
	// currentTime := time.Now()

	// session, err := sessionStore.Get(r, fmt.Sprintf("Session-%s", currentTime))

	session, err := sessionStore.Get(r, "Session-1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["me"] = 36

	err = sessions.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
