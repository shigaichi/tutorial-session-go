package model

type Category struct {
	CategoryID   int    `gorm:"primaryKey"`
	CategoryName string `gorm:"type:varchar(128)"`
}
