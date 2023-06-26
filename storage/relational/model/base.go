package model

// import (
// 	"time"

// 	"gorm.io/gorm"
// )

type BaseModel struct {
	ID        int64          `json:"id"`
	// CreatedAt time.Time      `json:"createdAt" swaggertype:"string"`
	// UpdatedAt time.Time      `json:"updatedAt" swaggertype:"string"`
	// DeletedAt gorm.DeletedAt `json:"deletedAt" swaggertype:"string"`
}

type UniqueCodeBase struct {
	Code        string `json:"code" label:"唯一标识" binding:"required,min=1,max=32"`
	Label       string `json:"label" label:"名称"  binding:"required,min=1,max=64"`
	Description string `json:"description" label:"描述" binding:"omitempty,max=255"`
}

type IDsQueryIn struct {
	IDs int64 `json:"id"`
}

type IDQueryIn struct {
	ID int64 `json:"id"`
}
