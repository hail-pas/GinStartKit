package model

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int64          `gorm:"primaryKey" json:"Version"`
	CreatedAt int64          `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
