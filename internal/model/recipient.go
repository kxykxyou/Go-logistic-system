package model

type Recipient struct {
	BaseModel
	Name    string `gorm:"not null; size: 30"`
	Address string `gorm:"not null; size: 100"`
	Phone   string `gorm:"not null; size: 10"`
	Order   []Order
}
