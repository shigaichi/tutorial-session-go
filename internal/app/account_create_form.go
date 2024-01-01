package app

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/shigaichi/tutorial-session-go/internal/domain/model"
)

type AccountCreateForm struct {
	Name            string
	Email           string
	Password        string
	ConfirmPassword string
	Birthday        time.Time
	Zip             string
	Address         string
}

func (f AccountCreateForm) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Name, validation.Required, validation.Length(1, 255)),
		validation.Field(&f.Email, validation.Required, validation.Length(1, 255)),
		validation.Field(&f.Password, validation.Required, validation.Length(4, 0)),
		validation.Field(&f.ConfirmPassword, validation.Required, validation.Length(4, 0)),
		validation.Field(&f.Zip, validation.Required, validation.Length(7, 7)),
		validation.Field(&f.Address, validation.Required, validation.Length(1, 255)),
	)
}

func (f AccountCreateForm) IsPasswordConfirmed() bool {
	return f.Password == f.ConfirmPassword
}

func (f AccountCreateForm) ToModel() model.Account {
	return model.Account{
		Email:    f.Email,
		Name:     f.Name,
		Birthday: f.Birthday,
		Zip:      f.Zip,
		Address:  f.Address,
	}
}
