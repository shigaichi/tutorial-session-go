package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Goods struct {
	GoodsID     string `gorm:"primaryKey;type:varchar(36)"`
	GoodsName   string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(512)"`
	CategoryID  int
	Price       int
	Category    Category `gorm:"foreignKey:CategoryID"`
}

func (g *Goods) BeforeCreate(_ *gorm.DB) error {
	random, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	g.GoodsID = random.String()

	return nil
}
