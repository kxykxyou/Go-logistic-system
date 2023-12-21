package model

type Product struct {
	BaseModel
	Name  string `gorm:"not null; size: 30"`
	Order []Order
}
