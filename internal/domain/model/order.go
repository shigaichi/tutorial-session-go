package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	OrderID   string `gorm:"primaryKey;type:varchar(36)"`
	Email     string `gorm:"type:varchar(255)"`
	OrderDate time.Time
}

func (g *Order) BeforeCreate(_ *gorm.DB) error {
	random, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	g.OrderID = random.String()

	return nil
}
