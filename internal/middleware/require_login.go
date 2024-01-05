package middleware

import (
	"net/http"

	log "github.com/shigaichi/tutorial-session-go/internal/logger"
	"go.uber.org/zap"
)

func RequireLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := SessionStore.Get(r, SessionName)
		if err != nil {
			log.Error("require_login", zap.Error(err))
			http.Error(w, "get session in RequireLogin middleware", http.StatusInternalServerError)
			return
		}

		email := session.Values["email"]

		if email == nil {
			http.Redirect(w, r, "/loginForm", http.StatusTemporaryRedirect)
		}
		next.ServeHTTP(w, r)
	})
}
