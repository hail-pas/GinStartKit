package api

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/api/v1/account"
	"github.com/hail-pas/GinStartKit/api/v1/auth"
	"github.com/hail-pas/GinStartKit/api/v1/requestRecord"
	"github.com/hail-pas/GinStartKit/api/v1/role"
	"github.com/hail-pas/GinStartKit/api/v1/system"
	systemReource "github.com/hail-pas/GinStartKit/api/v1/systemResource"
	assets "github.com/hail-pas/GinStartKit/asset"
	_ "github.com/hail-pas/GinStartKit/docs"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"html/template"
	"net/http"
	"time"
)

func RootEngine() *gin.Engine {
	rootRouter := gin.Default()
	// 模版挂载
	rootRouter.SetHTMLTemplate(
		template.Must(template.New("").
			ParseFS(assets.Templates, "templates/**/*"),
		), // 两级目录
	)

	err := rootRouter.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}
	middleware.RegisterMiddlewares(rootRouter)

	// pprof
	pprof.Register(rootRouter)

	// 文件路由
	{
		// 文件挂载
		fileRouter := rootRouter.Group("media")
		fileRouter.StaticFS("", http.Dir("./asset/files"))
	}

	v1Router := rootRouter.Group("/api/v1")

	// 无需登录路由
	{
		publicGroup := v1Router.Group("")
		// health
		publicGroup.GET("/health", func(c *gin.Context) {
			response.OkWithData(c, gin.H{
				"status":      "ok",
				"timestamp":   time.Now(),
				"environment": global.Configuration.System.Environment,
			})
		})
		// Auth
		auth.RegisterRouter(publicGroup)
		// system
		system.RegisterRouter(publicGroup)
		// swagger
		if global.Configuration.System.Debug {
			// => /api/v1/swagger/index.html
			publicGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}

	// 需要登录路由
	{
		privateGroup := v1Router.Group("")
		privateGroup.Use(
			middleware.AuthMiddlewareFunc(), middleware.PermissionChecker(),
		)
		// requestRecord
		{
			requestRecord.RegisterRouter(privateGroup)
		}
		// 请求记录路由
		{
			requestRecordGroup := privateGroup.Group("")
			requestRecordGroup.Use(
				middleware.RequestRecorder(),
			)
			// Account
			{
				account.RegisterRouter(requestRecordGroup)
			}
			// role
			{
				role.RegisterRouter(requestRecordGroup)
			}
			// systemResource
			{
				systemReource.RegisterRouter(requestRecordGroup)
			}
		}
	}

	return rootRouter
}
