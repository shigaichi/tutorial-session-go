package repository

import (
	"github.com/cockroachdb/errors"
	"github.com/shigaichi/tutorial-session-go/internal/domain/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() ([]model.Category, error)
}

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) CategoryRepositoryImpl {
	return CategoryRepositoryImpl{db: db}
}

func (c CategoryRepositoryImpl) FindAll() ([]model.Category, error) {
	var categories []model.Category
	err := c.db.Find(&categories).Error
	if err != nil {
		return nil, errors.Wrap(err, "FindAll")
	}
	return categories, nil
}
