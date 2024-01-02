package route

import (
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/shigaichi/tutorial-session-go/internal/app"
	"github.com/shigaichi/tutorial-session-go/internal/logger"
)

type Route interface {
	InitRouting() (*mux.Router, error)
}

type InitRoute struct {
	lh  app.LoginHandler
	ach app.AccountCreateHandler
}

func NewInitRoute(lh app.LoginHandler, ach app.AccountCreateHandler) InitRoute {
	return InitRoute{lh: lh, ach: ach}
}

func (i InitRoute) InitRouting() (*mux.Router, error) {
	r := mux.NewRouter()
	csrfMiddleware := csrf.Protect([]byte(os.Getenv("CSRF_KEY")))
	r.Use(csrfMiddleware)
	r.Use(handlers.RecoveryHandler(handlers.RecoveryLogger(&logger.ZapRecoveryLogger{Logger: logger.Logger})))

	fileServer := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
	r.HandleFunc("/loginForm", i.lh.LoginFormHandler).Methods("GET")
	r.HandleFunc("/authenticate", i.lh.AuthenticateHandler).Methods("POST")

	r.HandleFunc("/account/create", i.ach.ConfirmCreate).Methods("POST").Queries("confirm", "")
	r.HandleFunc("/account/create", i.ach.FinishCreate).Methods("GET").Queries("finish", "")
	r.HandleFunc("/account/create", i.ach.ShowCreateForm).Methods("GET")
	r.HandleFunc("/account/create", i.ach.Update).Methods("POST")
	return r, nil
}
