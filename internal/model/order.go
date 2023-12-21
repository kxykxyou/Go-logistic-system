package model

type Order struct {
	BaseModel
	RecipientId uint `gorm:"not null"`
	ProductId   uint `gorm:"not null"`
}
