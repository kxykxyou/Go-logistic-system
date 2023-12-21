package model

type BaseModel struct {
	id        uint `gorm:"column:Id;primaryKey;autoIncrement"`
	createdAt uint `gorm:"autoCreateTime:<-create"`
	updatedAt uint `gorm:"autoUpdateTime:<-"`
	deletedAt uint `gorm:"default:null"`
}
