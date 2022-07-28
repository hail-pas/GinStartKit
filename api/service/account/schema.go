package account

import "github.com/hail-pas/GinStartKit/storage/relational/model"

type UserResponseModel struct {
	model.BaseModel
	model.UsernameField
	model.PhoneField
	model.UserOtherInfo
	Systems []model.System `json:"systems"`
}
