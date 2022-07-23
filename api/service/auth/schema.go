package auth

import "github.com/hail-pas/GinStartKit/storage/relational/model"

type UserRegisterIn struct {
	model.UsernameField
	model.PhoneField
	model.UserOtherInfo
	model.PasswordField
}
