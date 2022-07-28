package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/api/service/auth"
)

// Register
// @Tags Auth
// @Summary 注册接口
// @accept application/json
// @Produce application/json
// @Param  path int true "ID"
// @Success 200 {object} response.response.Resp
// @Failure 200 {object} response.response.Resp
// @Router /api/auth/register [post]
func Register(c *gin.Context) {
	userRegisterIn, err := auth.ValidateRegisterIn(c)
	if err != nil {
		return
	}
	_ = auth.CreateUser(c, userRegisterIn)
}
