package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/global/constant"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
)

type UserRegisterIn struct {
	model.UsernameField
	model.PhoneField
	model.UserOtherInfo
	model.PasswordField
	model.SystemIDField
}

func ValidateRegisterIn(c *gin.Context) (*UserRegisterIn, error) {
	var userRegisterIn UserRegisterIn

	err := c.ShouldBindJSON(&userRegisterIn)
	if err != nil {
		response.ErrorResp(c, err)
		return nil, err
	}

	var exists bool

	err = global.RelationalDatabase.Table("system").Select("id").Where(
		"id = ?",
		userRegisterIn.SystemId,
	).Find(&exists).Error

	if err != nil || !exists {
		response.Response(
			c,
			constant.CodeContentNotFound,
			nil,
			fmt.Sprintf(constant.MessageContentNotFound, fmt.Sprintf("ID为%d的系统", userRegisterIn.SystemId)),
			-1,
			-1,
			-1,
		)
		return nil, global.BreakError
	}

	exists = false

	err = global.RelationalDatabase.Model(model.User{}).Select("id").Where(
		"username = ?",
		userRegisterIn.Username,
	).Find(&exists).Error

	if err != nil || exists {
		response.FailWithMessage(c, constant.UserWithUsernameExisted)
		return nil, global.BreakError
	}

	err = global.RelationalDatabase.Model(model.User{}).Select("id").Where(
		"phone = ?",
		userRegisterIn.Phone,
	).Find(&exists).Error

	if err != nil || exists {
		response.FailWithMessage(c, constant.UserWithPhoneExisted)
		return nil, global.BreakError
	}

	return &userRegisterIn, nil

}
