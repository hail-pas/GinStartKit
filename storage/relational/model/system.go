package model

type SystemSimpleFields struct {
	ID int64 `json:"id"`
	UniqueCodeBase
}

func (SystemSimpleFields) TableName() string {
	return "system"
}

type System struct {
	BaseModel
	UniqueCodeBase
	SystemResources []SystemResource `json:"systemResources" gorm:"many2many:system_with_system_resource;" label:"系统资源列表"`
	Roles           []Role           `json:"roles" label:"角色列表"`
	Users           []User           `json:"users" gorm:"many2many:system_with_user;joinForeignKey:SystemID;joinReferences:UserID" label:"用户列表"`
}

type SystemIDField struct {
	SystemId int64 `json:"systemId" binding:"required,min=1" label:"系统ID"`
}

type SystemWithUser struct {
	SystemIDField
	UserIDField
}
