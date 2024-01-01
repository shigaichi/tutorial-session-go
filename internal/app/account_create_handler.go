package app

import (
	"encoding/gob"
	"errors"
	"html/template"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/csrf"
	"github.com/shigaichi/tutorial-session-go/internal/domain/service"
	log "github.com/shigaichi/tutorial-session-go/internal/logger"
	"github.com/shigaichi/tutorial-session-go/internal/middleware"
	"go.uber.org/zap"
)

type AccountCreate interface {
	ShowCreateForm(w http.ResponseWriter, r *http.Request)
	ConfirmCreate(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	FinishCreate(w http.ResponseWriter, r *http.Request)
}

type AccountCreteHandler struct {
	as service.AccountService
}

func NewAccountCreteHandler(as service.AccountService) AccountCreteHandler {
	gob.Register(AccountCreateForm{})

	return AccountCreteHandler{as: as}
}

const sessionName = "accountCreateForm"

func (h AccountCreteHandler) ShowCreateForm(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/layout/template.gohtml", "templates/account/createForm.gohtml", "templates/layout/footer.gohtml"))

	if err := t.Execute(w, map[string]interface{}{csrf.TemplateTag: csrf.TemplateField(r), "title": "Account Create Page"}); err != nil {
		h.handleInternalError(w, err)
		return
	}
}

func (h AccountCreteHandler) ConfirmCreate(w http.ResponseWriter, r *http.Request) {
	store := middleware.SessionStore
	session, err := store.Get(r, middleware.SessionName)
	if err != nil {
		h.handleInternalError(w, err)
		return
	}

	f := AccountCreateForm{
		Name:            r.FormValue("name"),
		Email:           r.FormValue("email"),
		Password:        r.FormValue("password"),
		ConfirmPassword: r.FormValue("confirmPassword"),
		Zip:             r.FormValue("zip"),
		Address:         r.FormValue("address"),
	}

	birthday, err := time.Parse("2006-01-02", r.FormValue("birthday"))
	if err != nil {
		h.handleInternalError(w, err)
		return
	} else {
		f.Birthday = birthday
	}

	err = f.Validate()
	var valErr validation.Errors
	if err != nil && !errors.As(err, &valErr) {
		h.handleInternalError(w, err)
		return
	} else if err != nil {
		t := template.Must(template.ParseFiles("templates/layout/template.gohtml", "templates/account/createForm.gohtml", "templates/layout/footer.gohtml"))

		if err := t.Execute(w, map[string]interface{}{csrf.TemplateTag: csrf.TemplateField(r), "title": "Account Create Page", "errors": valErr}); err != nil {
			h.handleInternalError(w, err)
			return
		}

		return
	}

	if !f.IsPasswordConfirmed() {
		t := template.Must(template.ParseFiles("templates/layout/template.gohtml", "templates/account/createForm.gohtml", "templates/layout/footer.gohtml"))

		if err := t.Execute(w, map[string]interface{}{csrf.TemplateTag: csrf.TemplateField(r), "title": "Account Create Page", "errors": map[string]string{"Password": "Password confirmation is failed."}}); err != nil {
			h.handleInternalError(w, err)
			return
		}

		return
	}

	session.Values[sessionName] = f
	err = session.Save(r, w)
	if err != nil {
		h.handleInternalError(w, err)
		return
	}

	t := template.Must(template.ParseFiles("templates/layout/template.gohtml", "templates/account/createConfirm.gohtml", "templates/layout/footer.gohtml"))
	if err := t.Execute(w, map[string]interface{}{csrf.TemplateTag: csrf.TemplateField(r), "title": "Item List Page", "accountCreateForm": f}); err != nil {
		h.handleInternalError(w, err)
		return
	}
}

func (h AccountCreteHandler) Update(w http.ResponseWriter, r *http.Request) {
	if _, ok := r.Form["redoForm"]; ok {
		h.ShowCreateForm(w, r)
		return
	}

	// TODO: implement create account
}

func (h AccountCreteHandler) FinishCreate(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h AccountCreteHandler) handleInternalError(w http.ResponseWriter, err error) {
	log.Error("error in account_create_handler", zap.Error(err))
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
