package model

type BaseModel struct {
	ID        uint  `gorm:"primarykey" json:"ID"`
	CreatedAt int64 `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}
