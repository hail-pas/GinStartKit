package model

import (
	"github.com/google/uuid"
)

type PasswordField struct {
	Password string `json:"-"  gorm:"comment:用户登录密码"` // 用户登录密码
}

type PhoneField struct {
	Phone string `json:"phone"  gorm:"uniqueIndex;comment:用户手机号"` // 用户手机号
}

type UsernameField struct {
	Username string `json:"userName" gorm:"uniqueIndex;comment:用户登录名"` // 用户登录名
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
	UUID     uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`                // 用户UUID
	NickName string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`       // 用户昵称
	Avatar   string    `json:"avatar" gorm:"comment:用户头像"`                      // 用户头像
	Email    string    `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enabled  bool      `json:"enable" gorm:"default:1;comment:用户是否被禁用 1启用 0禁用"` //用户是否被禁用 1启用 0禁用
}

type User struct {
	BaseModel
	UsernameField
	PhoneField
	PasswordField
	UserOtherInfo
}
