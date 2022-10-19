package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/fvbock/endless"
	"github.com/hail-pas/GinStartKit/api"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/initialize"
	"github.com/rs/zerolog/log"
)

func initializeGlobal(configPath string) {
	// 初始化操作
	initialize.Configuration(configPath)
	initialize.Redis()
	initialize.GormDB()
	initialize.ValidateWithTranslation("zh")
}

// @title 			GInStartKit
// @version 		1.0
// @description 	Start Kit of Gin
// @termsOfService	https://github.com/hail-pas/GinStartKit

// @license.name	MIT
// @license.url 	https://github.com/hail-pas/GinStartKit

// @contact.name   hypo-fiasco
// @contact.url    https://github.com/hail-pas/GinStartKit
// @contact.email  hypofiasco@gmail.com

// @host 		127.0.0.1:8000
// @BasePath 	/api/v1/
// @schemes 	http

// @securityDefinitions.apikey  Jwt
// @in                          header
// @name                        Authorization
func main() {
	configFile := flag.String("conf", "", "Path to the configuration file")
	flag.Parse()

	dir, _ := os.Getwd()

	fmt.Print(dir)

	if *configFile == "" {
		environment := os.Getenv("environment")
		if environment != "" {
			*configFile = fmt.Sprintf("config/content/%s.yaml", environment)
		} else {
			*configFile = "config/content/default.yaml"
		}
	}

	initializeGlobal(*configFile)

	log.Info().Msgf("%+v", global.Configuration)

	engine := api.RootEngine()

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

	err := s.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
