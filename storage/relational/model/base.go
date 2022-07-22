package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int64          `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
