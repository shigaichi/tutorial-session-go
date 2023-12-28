package model

type OrderLine struct {
	OrderID  string `gorm:"primaryKey;type:varchar(36)"`
	GoodsID  string `gorm:"primaryKey;type:varchar(36)"`
	Quantity int
	Order    Order `gorm:"foreignKey:OrderID"`
	Goods    Goods `gorm:"foreignKey:GoodsID"`
}
