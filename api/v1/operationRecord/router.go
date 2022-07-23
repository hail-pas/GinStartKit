package operationRecord

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(parentRouter *gin.RouterGroup) {
	router := parentRouter.Group("/operationRecord")
	router.GET("", List)
}
