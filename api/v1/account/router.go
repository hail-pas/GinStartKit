package account

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/middleware"
)

func RegisterRouter(parentRouter *gin.RouterGroup) {
	router := parentRouter.Group("/account")
	router.GET("", List)
	router.GET("/token/refresh", middleware.GetAuthJwtMiddleware().RefreshHandler)
}
