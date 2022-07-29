package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/api/service/auth"
	_ "github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/middleware"
	_ "github.com/hail-pas/GinStartKit/storage/relational/model"
)

// Register
// @Tags Auth
// @Summary 注册接口
// @accept json
// @Produce json
// @Param data body auth.UserRegisterIn true "注册body"
// @Success 200 {object} response.Resp 成功
// @Router /auth/register [post]
func Register(c *gin.Context) {
	userRegisterIn, err := auth.ValidateRegisterIn(c)
	if err != nil {
		return
	}
	_ = auth.CreateUser(c, userRegisterIn)
}

// Login
// @Tags Auth
// @Summary 登录接口
// @accept json
// @Produce json
// @Param data body model.UserLoginWithPhone true "登录body"
// @Success 200 {object} response.Resp{data=model.User} 成功
// @Router /auth/login [post]
func Login(c *gin.Context) {
	middleware.GetAuthJwtMiddleware().LoginHandler(c)
}
