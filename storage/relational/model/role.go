package model

type Role struct {
	BaseModel
	UniqueCodeBase
	SystemIDField
	System          System           `json:"system"`
	SystemResources []SystemResource `json:"systemResources" gorm:"many2many:role_with_system_resource;" label:"系统资源列表"`
}
