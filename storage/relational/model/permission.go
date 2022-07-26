package model

type Permission struct {
	BaseModel
	UniqueCodeBase
	SystemResources []SystemResource `json:"systemResources" gorm:"many2many:system_resource_with_permission;" label:"系统资源列表"`
}
