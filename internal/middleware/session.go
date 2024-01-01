package middleware

import (
	"encoding/gob"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

func init() {
	SessionStore.Options.HttpOnly = true
	SessionStore.MaxAge(86400 * 5)
	SessionStore.Options.SameSite = http.SameSiteLaxMode
	SessionStore.Options.Secure = true
	gob.Register(map[string]string{})
}

var SessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

const SessionName = "tutorial-session-go"
