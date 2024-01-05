package repository

import (
	"github.com/cockroachdb/errors"
	"github.com/shigaichi/tutorial-session-go/internal/domain/model"
	util "github.com/shigaichi/tutorial-session-go/internal/orm"
	"gorm.io/gorm"
)

type GoodsRepository interface {
	CountByCategoryId(categoryId int) (int64, error)
	FindPageByCategoryId(categoryId, pageNumber, pageSize int) ([]model.Goods, error)
}

type GoodsRepositoryImpl struct {
	db *gorm.DB
}

func (i GoodsRepositoryImpl) CountByCategoryId(categoryId int) (int64, error) {
	var count int64
	err := i.db.Model(&model.Goods{}).Where("category_id = ?", categoryId).Count(&count).Error
	if err != nil {
		return 0, errors.Wrap(err, "CountByCategoryId")
	}
	return count, nil
}

func NewGoodsRepositoryImpl(db *gorm.DB) GoodsRepositoryImpl {
	return GoodsRepositoryImpl{db: db}
}

func (i GoodsRepositoryImpl) FindPageByCategoryId(categoryId, pageNumber, pageSize int) ([]model.Goods, error) {
	var goods []model.Goods
	err := i.db.Scopes(util.Paginate(pageNumber, pageSize)).Where("category_id = ?", categoryId).Order("goods_id").Take(&goods).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "take by categoryId")
	}

	return goods, nil
}
