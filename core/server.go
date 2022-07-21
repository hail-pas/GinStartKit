package main

import (
	"flag"
	"fmt"
	"github.com/fvbock/endless"
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
	initialize.ValidateWithTranslation()
}

func main() {
	configFile := flag.String("conf", "./config/content/default.yaml", "Path to the configuration file")
	flag.Parse()

	initializeGlobal(*configFile)

	log.Info().Msgf("%+v", global.Configuration)

	engine := router.RootEngine()

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
