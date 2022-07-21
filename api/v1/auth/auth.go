package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/global/constant"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
	"github.com/jinzhu/copier"
)

type UserRegisterIn struct {
	model.UsernameField
	model.PhoneField
	model.PasswordField
	model.UserOtherInfo
}

func (receiver UserRegisterIn) validate() error {
	validate := validator.New()
	return validate.Struct(receiver)
}

// Register
// @Tags Auth
// @Summary 注册接口
// @accept application/json
// @Produce application/json
// @Router /api/auth/register [post]
func Register(c *gin.Context) {

	var userRegisterIn UserRegisterIn

	err := c.BindJSON(&userRegisterIn)
	if err != nil {
		response.BadRequest(c, constant.ParamBindError)
		return
	}

	err = userRegisterIn.validate()
	if err != nil {
		response.BadRequest(c, constant.ParamCheckError)
		return
	}

	var user model.User

	_ = copier.Copy(&user, &userRegisterIn)
	user.UUID, _ = uuid.NewUUID()
	global.RelationalDatabase.Create(&user)
	response.Ok(c)
}
