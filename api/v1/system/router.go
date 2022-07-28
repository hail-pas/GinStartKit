package system

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(parentRouter *gin.RouterGroup) {
	router := parentRouter.Group("/system")
	router.GET("", func(context *gin.Context) {

	})
}
