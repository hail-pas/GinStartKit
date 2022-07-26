package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hail-pas/GinStartKit/api/service/auth"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/global/common/utils"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
	"github.com/jinzhu/copier"
)

// Register
// @Tags Auth
// @Summary 注册接口
// @accept application/json
// @Produce application/json
// @Router /api/auth/register [post]
func Register(c *gin.Context) {
	userRegisterIn, err := auth.ValidateRegisterIn(c)
	if err != nil {
		return
	}
	var user model.User
	_ = copier.Copy(&user, userRegisterIn)
	user.UUID, _ = uuid.NewUUID()
	user.Enabled = true
	user.Password = utils.PasswordHash(user.Password)
	global.RelationalDatabase.Create(&user)
	response.Ok(c)
}
