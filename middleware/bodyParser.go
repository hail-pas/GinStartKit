package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func BodyParser() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info().Msg("Body Parsing")
		//c.Next()
	}
}
