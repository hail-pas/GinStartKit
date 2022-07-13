package initialize

import (
	"github.com/hail-pas/GinStartKit/config"
	"github.com/hail-pas/GinStartKit/global"
)

func Configuration(configFilePath string) {
	if global.Configuration == nil {
		if configFilePath != "" {
			global.Configuration = config.SetConfig(configFilePath)
		} else {
			global.Configuration = config.SetConfig("")
		}
	}
}
