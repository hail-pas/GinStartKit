package main

import (
	"flag"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/api/router"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/initialize"
	"github.com/rs/zerolog/log"
	"time"
)

func initializeGlobal(configPath string) {
	// 初始化操作
	initialize.Configuration(configPath)
	initialize.Redis()
	initialize.GormDB()
}

func registerMiddlewares(r *gin.Engine) {
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

	r.Use()
}

func main() {
	configFile := flag.String("conf", "./config/content/default.yaml", "Path to the configuration file")
	flag.Parse()

	initializeGlobal(*configFile)

	log.Info().Msgf("%+v", global.Configuration)

	engine := router.RootEngine()

	err := engine.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}

	registerMiddlewares(engine)

	serverAddress := fmt.Sprintf(
		"%s:%d",
		global.Configuration.System.TcpAddr.IP.String(),
		global.Configuration.System.TcpAddr.Port,
	)

	s := endless.NewServer(serverAddress, engine)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20

	time.Sleep(10 * time.Microsecond)
	fmt.Println("=============================")
	fmt.Printf("Start Running on %v\n", global.Configuration.System.TcpAddr.String())
	fmt.Println("=============================")

	err = s.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
