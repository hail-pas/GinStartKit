package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/global"
	"time"
)

func RegisterMiddlewares(r *gin.Engine) {
	if global.Configuration.System.CorsConfig.AllowAll {
		r.Use(cors.Default())
	} else {
		r.Use(cors.New(cors.Config{
			AllowOrigins:     global.Configuration.System.CorsConfig.AllowOrigins,
			AllowMethods:     global.Configuration.System.CorsConfig.AllowMethods,
			AllowHeaders:     global.Configuration.System.CorsConfig.AllowHeaders,
			ExposeHeaders:    global.Configuration.System.CorsConfig.ExposeHeaders,
			AllowCredentials: global.Configuration.System.CorsConfig.AllowCredentials,
			MaxAge:           time.Duration(global.Configuration.System.CorsConfig.MaxAge) * time.Second,
		}))
	}
}
