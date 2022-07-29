package account

import "github.com/hail-pas/GinStartKit/storage/relational/model"

type UserResponseModel struct {
	model.BaseModel
	model.UsernameField
	model.PhoneField
	model.UserOtherInfo
	Systems []model.SystemSimpleFields `json:"systems" gorm:"many2many:system_with_user;joinForeignKey:UserID;joinReferences:SystemID" label:"系统列表"`
}
