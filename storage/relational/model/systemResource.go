package model

type SystemResource struct {
	BaseModel
	UniqueCodeBase
	ParentId       int64           `json:"parentId" label:"父节点ID"`
	Parent         *SystemResource `json:"parent"  label:"父节点" gorm:"foreignKey:ParentId"`
	ReferenceToId  int64           `json:"referenceToId" label:"跳转节点ID"`
	ReferenceTo    *SystemResource `json:"referenceTo" label:"跳转节点" gorm:"foreignKey:ReferenceToId"`
	FrontRoutePath string          `json:"frontRoutePath" label:"前端路由"`
	IconPath       string          `json:"iconPath" label:"图标地址"`
	Type           string          `json:"type"`
	OrderNum       int             `json:"orderNum"`
	Enabled        bool            `json:"enabled"`
	Permissions    []Permission    `json:"systemResources" gorm:"many2many:system_resource_with_permission;" label:"权限列表"`
	Roles          []Role          `json:"roles" gorm:"many2many:role_with_system_resource;" label:"角色列表"`
}
