package service

import (
	"github.com/cockroachdb/errors"
	"github.com/shigaichi/tutorial-session-go/internal/domain/model"
	"github.com/shigaichi/tutorial-session-go/internal/domain/repository"
)

type CategoryService interface {
	FindAll() ([]model.Category, error)
}

type CategoryServiceImpl struct {
	cr repository.CategoryRepository
}

func NewCategoryServiceImpl(cr repository.CategoryRepository) CategoryServiceImpl {
	return CategoryServiceImpl{cr: cr}
}

func (i CategoryServiceImpl) FindAll() ([]model.Category, error) {
	categories, err := i.cr.FindAll()
	if err != nil {
		return nil, errors.Wrap(err, "FindAll")
	}

	return categories, nil
}
