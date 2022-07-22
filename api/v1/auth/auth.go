package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/global/common/utils"
	"github.com/hail-pas/GinStartKit/global/constant"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
	"github.com/jinzhu/copier"
)

type UserRegisterIn struct {
	model.UsernameField
	model.PhoneField
	model.UserOtherInfo
	model.PasswordField
}

// Register
// @Tags Auth
// @Summary 注册接口
// @accept application/json
// @Produce application/json
// @Router /api/auth/register [post]
func Register(c *gin.Context) {

	var userRegisterIn UserRegisterIn

	err := c.ShouldBindJSON(&userRegisterIn)
	if err != nil {
		response.ErrorResp(c, err)
		return
	}

	var exists bool

	err = global.RelationalDatabase.Model(model.User{}).Select("id").Where(
		"username = ?",
		userRegisterIn.Username,
	).Find(&exists).Error

	if err != nil || exists {
		response.FailWithMessage(c, constant.UserWithUsernameExisted)
		return
	}

	err = global.RelationalDatabase.Model(model.User{}).Select("id").Where(
		"phone = ?",
		userRegisterIn.Phone,
	).Find(&exists).Error

	if err != nil || exists {
		response.FailWithMessage(c, constant.UserWithPhoneExisted)
		return
	}

	var user model.User
	_ = copier.Copy(&user, &userRegisterIn)
	user.UUID, _ = uuid.NewUUID()
	user.Enabled = true
	user.Password = utils.PasswordHash(user.Password)
	global.RelationalDatabase.Create(&user)
	response.Ok(c)
}
