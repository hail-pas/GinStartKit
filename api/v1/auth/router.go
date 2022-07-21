package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/middleware"
	"github.com/hail-pas/GinStartKit/pkg/utils"
)

func RegisterRouter(parentRouter *gin.RouterGroup) {
	router := parentRouter.Group("/auth").Use(middleware.OperationRecord())
	router.POST("/register", Register)
	router.POST("/login", utils.GetAuthJwtMiddleware().LoginHandler)
	router.GET("/token/refresh", utils.GetAuthJwtMiddleware().RefreshHandler)
}
