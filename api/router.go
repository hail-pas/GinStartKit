package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/api/v1/account"
	"github.com/hail-pas/GinStartKit/api/v1/auth"
	"github.com/hail-pas/GinStartKit/api/v1/requestRecord"
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
	v1Router := rootRouter.Group("/api/v1")
	PublicGroup := v1Router.Group("")
	{
		// health
		PublicGroup.GET("/health", func(c *gin.Context) {
			response.Ok(c)
		})
		PublicGroup.GET("/health1", func(c *gin.Context) {
			response.Ok(c)
		})
		// Auth
		auth.RegisterRouter(PublicGroup)
	}
	PrivateGroup := v1Router.Group("")
	PrivateGroup.Use(
		middleware.AuthMiddlewareFunc(), middleware.PermissionChecker(),
	)
	requestRecordGroup := PrivateGroup.Group("")
	requestRecordGroup.Use(
		middleware.RequestRecorder(),
	)
	// Account
	{
		account.RegisterRouter(requestRecordGroup)
	}
	// requestRecord
	{
		requestRecord.RegisterRouter(PrivateGroup)
	}

	return rootRouter
}
