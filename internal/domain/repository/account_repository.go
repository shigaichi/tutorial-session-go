package repository

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/shigaichi/tutorial-session-go/internal/domain/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	FindByEmail(ctx context.Context, email string) (model.Account, error)
}

type AccountRepositoryImpl struct {
	db *gorm.DB
}

func NewAccountRepositoryImpl(db *gorm.DB) AccountRepositoryImpl {
	return AccountRepositoryImpl{db: db}
}

func (i AccountRepositoryImpl) FindByEmail(ctx context.Context, email string) (model.Account, error) {
	a := model.Account{}
	err := i.db.WithContext(ctx).Where("email = ?", email).Take(&a).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Account{}, nil
		}
		return model.Account{}, errors.Wrap(err, "failed to find by email")
	}
	return a, nil
}
