package model

import "time"

type LogisticDetail struct {
	BaseModel
	OrderId    uint `gorm:"not null"`
	LocationId uint `gorm:"not null"`
	Date       time.Time
	Status     string `gorm:"size: 30"`
	Order      Order
	Location   Location
}
