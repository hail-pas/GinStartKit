package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/api/schema/common/response"
	"github.com/hail-pas/GinStartKit/api/v1/account"
	"github.com/hail-pas/GinStartKit/api/v1/auth"
	"github.com/hail-pas/GinStartKit/middleware"
)

func RootEngine() *gin.Engine {
	rootRouter := gin.Default()
	err := rootRouter.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}
	middleware.RegisterMiddlewares(rootRouter)
	PublicGroup := rootRouter.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			response.Ok(c)
		})
	}
	PrivateGroup := rootRouter.Group("/api/v1")

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
