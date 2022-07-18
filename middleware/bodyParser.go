package middleware

import (
	"github.com/gin-gonic/gin"
)

func BodyParser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//log.Info().Msg("Body Parsing")
		//response.FailWithMessage(c, "解析Body失败")
		//c.Next()
	}
}
