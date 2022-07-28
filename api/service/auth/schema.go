package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/global/common/utils"
	"github.com/hail-pas/GinStartKit/global/constant"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
)

type UserRegisterIn struct {
	model.UsernameField
	model.PhoneField
	model.UserOtherInfo
	model.PasswordField
	SystemIds []int64 `json:"systemIds" binding:"gt=0,dive,min=1,required" label:"系统IDs"`
}

func ValidateRegisterIn(c *gin.Context) (*UserRegisterIn, error) {
	var userRegisterIn UserRegisterIn

	err := c.ShouldBindJSON(&userRegisterIn)
	if err != nil {
		response.ErrorResp(c, err)
		return nil, err
	}

	userRegisterIn.SystemIds = new(utils.Set[int64]).Set(userRegisterIn.SystemIds...).Array()

	existed := true
	for _, systemId := range userRegisterIn.SystemIds {
		err = global.RelationalDatabase.Model(model.System{}).Select("count(*) > 0").Where(
			"id = ?",
			systemId,
		).Find(&existed).Error

		if err != nil || !existed {
			response.WithoutPageInfo[any](
				c,
				constant.CodeContentNotFound,
				nil,
				fmt.Sprintf(constant.MessageContentNotFound, fmt.Sprintf("ID为%d的系统", systemId)),
			)
			return nil, global.BreakError
		}
	}

	existed = false

	err = global.RelationalDatabase.Model(model.User{}).Select("count(*) > 0").Where(
		"username = ?",
		userRegisterIn.Username,
	).Find(&existed).Error

	if err != nil || existed {
		response.FailWithMessage(c, constant.UserWithUsernameExisted)
		return nil, global.BreakError
	}

	err = global.RelationalDatabase.Model(model.User{}).Select("count(*) > 0").Where(
		"phone = ?",
		userRegisterIn.Phone,
	).Find(&existed).Error

	if err != nil || existed {
		response.FailWithMessage(c, constant.UserWithPhoneExisted)
		return nil, global.BreakError
	}

	return &userRegisterIn, nil

}
