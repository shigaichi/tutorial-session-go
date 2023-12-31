package service

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/shigaichi/tutorial-session-go/internal/domain/model"
	"github.com/shigaichi/tutorial-session-go/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type AccountService interface {
	Authenticate(ctx context.Context, email string, password string) (model.Account, error)
}

type AccountServiceImpl struct {
	ar repository.AccountRepository
}

func NewAccountServiceImpl(a repository.AccountRepository) AccountServiceImpl {
	return AccountServiceImpl{ar: a}
}

func (i AccountServiceImpl) Authenticate(ctx context.Context, email string, password string) (model.Account, error) {
	account, err := i.ar.FindByEmail(ctx, email)
	if err != nil {
		return model.Account{}, errors.Wrap(err, "failed to find by email")
	}
	if account.ID == "" {
		return model.Account{}, nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.EncodedPassword), []byte(password))
	if err != nil {
		return model.Account{}, nil
	}
	return account, nil
}
