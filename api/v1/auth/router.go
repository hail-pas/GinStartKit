package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(parentRouter *gin.RouterGroup) {
	router := parentRouter.Group("/auth")
	router.POST("/register", Register)
	router.POST("/login", Login)
}
