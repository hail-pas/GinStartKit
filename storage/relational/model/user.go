package model

import "github.com/google/uuid"

type PasswordField struct {
	Password string `json:"-" binding:"min=10,max=20"` // 用户登录密码
}

type PhoneField struct {
	Phone string `json:"phone" binding:"required,len=11"` // 用户手机号
}

type UsernameField struct {
	Username string `json:"userName" binding:"required,min=8,max=32"` // 用户登录名
}

type UserLoginWithPhone struct {
	PhoneField
	PasswordField
}

type UserLoginWithUsername struct {
	UsernameField
	PasswordField
}

type UserOtherInfo struct {
	Nickname string `json:"nickname" binding:"required,max=64"` // 用户昵称
	Avatar   string `json:"avatar" binding:"url,max=256"`       // 用户头像
	Email    string `json:"email" binding:"email"`              // 用户邮箱
	Enabled  bool   `json:"enabled"`                            //用户是否被禁用 1启用 0禁用
}

type UserSelfGenerateFields struct {
	UUID uuid.UUID `json:"uuid" gorm:"column:uuid"` // 用户UUID
}

type User struct {
	BaseModel
	UsernameField
	PhoneField
	PasswordField
	UserOtherInfo
	UserSelfGenerateFields
}
