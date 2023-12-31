package app

import (
	"html/template"
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/shigaichi/tutorial-session-go/internal/domain/service"
	log "github.com/shigaichi/tutorial-session-go/internal/logger"
	"github.com/shigaichi/tutorial-session-go/internal/middleware"
	"go.uber.org/zap"
)

type Login interface {
	AuthenticateHandler(w http.ResponseWriter, r *http.Request)
	LoginFormHandler(w http.ResponseWriter, r *http.Request)
	LogoutHandler(w http.ResponseWriter, r *http.Request)
}

type LoginHandler struct {
	as service.AccountService
}

func NewLoginHandler(as service.AccountService) LoginHandler {
	return LoginHandler{as: as}
}

func (h LoginHandler) AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	store := middleware.SessionStore
	session, err := store.Get(r, middleware.SessionName)
	if err != nil {
		h.handleInternalError(w, err)
		return
	}

	email, password := r.FormValue("email"), r.FormValue("password")
	if email == "" || password == "" {
		if err = h.handleInvalidLogin(w, r, session, "email and password can not be empty."); err != nil {
			h.handleInternalError(w, err)
		}
		return
	}

	account, err := h.as.Authenticate(r.Context(), email, password)
	if err != nil {
		h.handleInternalError(w, err)
		return
	}

	if account.ID == "" {
		if err = h.handleInvalidLogin(w, r, session, "wrong email or password"); err != nil {
			h.handleInternalError(w, err)
		}
		return
	}

	session.Values["email"] = account.Email
	if err = session.Save(r, w); err != nil {
		h.handleInternalError(w, err)
		return
	}
	http.Redirect(w, r, "/goods", http.StatusFound)
}

func (h LoginHandler) handleInvalidLogin(w http.ResponseWriter, r *http.Request, session *sessions.Session, message string) error {
	session.AddFlash(message, "loginError")
	err := session.Save(r, w)
	if err != nil {
		return errors.Wrap(err, "failed to save flash")
	}
	http.Redirect(w, r, "/loginForm?error", http.StatusFound)
	return nil
}

func (h LoginHandler) LoginFormHandler(w http.ResponseWriter, r *http.Request) {
	store := middleware.SessionStore
	session, err := store.Get(r, middleware.SessionName)
	if err != nil {
		h.handleInternalError(w, err)
		return
	}

	f := session.Flashes("loginError")
	if err = session.Save(r, w); err != nil {
		h.handleInternalError(w, err)
		return
	}

	var loginError string
	if len(f) > 0 {
		loginError = f[0].(string)
	} else {
		loginError = ""
	}

	t := template.Must(template.ParseFiles("templates/layout/template.gohtml", "templates/login/loginForm.gohtml", "templates/layout/footer.gohtml"))
	if err = t.Execute(w, map[string]interface{}{csrf.TemplateTag: csrf.TemplateField(r), "loginError": loginError, "title": "Login Page"}); err != nil {
		h.handleInternalError(w, err)
		return
	}
}

func (h LoginHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h LoginHandler) handleInternalError(w http.ResponseWriter, err error) {
	log.Error("error in login_handler", zap.Error(err))
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
