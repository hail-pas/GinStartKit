package model

import "github.com/google/uuid"

type PasswordField struct {
	Password string `json:"password" binding:"min=10,max=20" label:"密码"` // 用户登录密码
}

type PhoneField struct {
	Phone string `json:"phone" binding:"required,len=11" label:"手机号"` // 用户手机号
}

type UsernameField struct {
	Username string `json:"username" binding:"required,min=6,max=32" label:"用户名"` // 用户登录名
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
	Nickname string `json:"nickname" binding:"required,max=64" label:"昵称"`        // 用户昵称
	Avatar   string `json:"avatar" binding:"omitempty,url,max=256" label:"头像URL"` // 用户头像
	Email    string `json:"email" binding:"omitempty,email" label:"邮箱"`           // 用户邮箱
	Enabled  bool   `json:"enabled" label:"启用状态"`                                 //用户是否被禁用 1启用 0禁用
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
