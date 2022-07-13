package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/core/middleware"
)

func RegisterRouter(parentRouter *gin.RouterGroup) {
	router := parentRouter.Group("/auth").Use(middleware.OperationRecord())
	router.GET("register", Register)
}
