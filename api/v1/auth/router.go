package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/middleware"
)

func RegisterRouter(parentRouter *gin.RouterGroup) {
	router := parentRouter.Group("/auth")
	router.POST("/register", Register)
	router.POST("/login", middleware.GetAuthJwtMiddleware().LoginHandler)
}
