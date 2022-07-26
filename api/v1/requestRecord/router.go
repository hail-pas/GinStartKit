package requestRecord

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(parentRouter *gin.RouterGroup) {
	router := parentRouter.Group("/requestRecord")
	router.GET("", List)
}
