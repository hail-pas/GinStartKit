package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/api/schema/common/response"
)

// Register
// @Tags Auth
// @Summary 注册接口
// @accept application/json
// @Produce application/json
// @Router /api/auth/register [post]
func Register(c *gin.Context) {

	response.Ok(c)
}
