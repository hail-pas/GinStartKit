package account

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/global/common/response"
)

func List(c *gin.Context) {
	response.OkWithPageData(c, make([]map[string]interface{}, 12), 10, 1, 129)
}
