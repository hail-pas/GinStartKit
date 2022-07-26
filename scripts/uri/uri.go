package main

import (
	"flag"
	"fmt"
	"github.com/hail-pas/GinStartKit/api"
	"github.com/hail-pas/GinStartKit/global/initialize"
)

func main() {
	configFile := flag.String("conf", "./config/content/default.yaml", "Path to the configuration file")
	flag.Parse()

	initialize.Configuration(*configFile)

	engine := api.RootEngine()
	routers := engine.Routes()
	for _, v := range routers {
		fmt.Print(v)
	}
}
