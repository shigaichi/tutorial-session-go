package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID                 string `dto:"primaryKey;size:36"`
	Email              string `gorm:"unique;not null;type:varchar(255)"`
	EncodedPassword    string `gorm:"type:varchar(255)"`
	Name               string `gorm:"type:varchar(255)"`
	Birthday           time.Time
	Zip                string `gorm:"type:char(7)"`
	Address            string `gorm:"type:varchar(255)"`
	CardNumber         string `gorm:"type:varchar(16)"`
	CardExpirationDate time.Time
	CardSecurityCode   string `gorm:"type:varchar(10)"`
}

func (a *Account) BeforeCreate(_ *gorm.DB) error {
	random, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	a.ID = random.String()

	return nil
}
