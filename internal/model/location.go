package model

type Location struct {
	BaseModel
	Title   string `gorm:"not null;size: 50"`
	City    string `gorm:"not null;size: 50"`
	Address string `gorm:"not null;size: 100"`
}
