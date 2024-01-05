package service

import (
	"github.com/cockroachdb/errors"
	"github.com/shigaichi/tutorial-session-go/internal/domain/model"
	"github.com/shigaichi/tutorial-session-go/internal/domain/repository"
)

type GoodsService interface {
	FindByCategoryId(categoryId, pageNumber, pageSize int) ([]model.Goods, int64, error)
}

type GoodsServiceImpl struct {
	gr repository.GoodsRepository
}

func NewGoodsServiceImpl(gr repository.GoodsRepository) GoodsServiceImpl {
	return GoodsServiceImpl{gr: gr}
}

func (i GoodsServiceImpl) FindByCategoryId(categoryId, pageNumber, pageSize int) ([]model.Goods, int64, error) {
	count, err := i.gr.CountByCategoryId(categoryId)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "CountByCategoryId: %d", categoryId)
	}

	if count == 0 {
		return nil, 0, nil
	}

	goods, err := i.gr.FindPageByCategoryId(categoryId, pageNumber, pageSize)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "find by categoryId categoryId: %d, pageNumber: %d, pageSize: %d", categoryId, pageNumber, pageSize)
	}

	return goods, count, nil
}
