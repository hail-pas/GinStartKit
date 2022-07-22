package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/api/v1/account"
	"github.com/hail-pas/GinStartKit/api/v1/auth"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/middleware"
)

func RootEngine() *gin.Engine {
	rootRouter := gin.Default()
	err := rootRouter.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}
	middleware.RegisterMiddlewares(rootRouter)
	vqRouter := rootRouter.Group("/api/v1")
	PublicGroup := vqRouter.Group("")
	{
		// health
		PublicGroup.GET("/health", func(c *gin.Context) {
			response.Ok(c)
		})
		// Auth
		auth.RegisterRouter(PublicGroup)
	}
	PrivateGroup := vqRouter.Group("")
	PrivateGroup.Use(
		middleware.AuthMiddlewareFunc(),
		middleware.OperationRecord(),
	)
	// Account
	{
		account.RegisterRouter(PrivateGroup)
	}

	return rootRouter
}
