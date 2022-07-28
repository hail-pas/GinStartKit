package role

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(parentRouter *gin.RouterGroup) {
	router := parentRouter.Group("/role")
	router.GET("", func(context *gin.Context) {
	})
}
