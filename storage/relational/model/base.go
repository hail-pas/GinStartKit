package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"Version"`
	CreatedAt sql.NullTime   `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
