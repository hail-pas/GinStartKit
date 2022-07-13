package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/api/schema/common/response"
	"github.com/hail-pas/GinStartKit/api/v1/account"
	"github.com/hail-pas/GinStartKit/api/v1/auth"
)

func RootEngine() *gin.Engine {
	rootRouter := gin.Default()
	PublicGroup := rootRouter.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			response.Ok(c)
			c.Abort()
		})
	}
	PrivateGroup := rootRouter.Group("/api")
	PrivateGroup.Use()

	// Auth
	{
		auth.RegisterRouter(PrivateGroup)
	}

	// Account
	{
		account.RegisterRouter(PrivateGroup)
	}

	return rootRouter
}
