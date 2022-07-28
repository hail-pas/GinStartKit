package systemReource

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(parentRouter *gin.RouterGroup) {
	router := parentRouter.Group("/systemResource")
	router.GET("", func(context *gin.Context) {

	})
}
