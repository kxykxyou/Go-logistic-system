package model

type BaseModel struct {
	Id        uint `gorm:"column:Id;primaryKey;autoIncrement"`
	CreatedAt uint `gorm:"autoCreateTime:<-create"`
	UpdatedAt uint `gorm:"autoUpdateTime:<-"`
	DeletedAt uint `gorm:"default:null"`
}
