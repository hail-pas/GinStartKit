package global

import (
	"context"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/hail-pas/GinStartKit/config"
	"github.com/hail-pas/GinStartKit/global/constant"
	"gorm.io/gorm"
)

var (
	Configuration      *config.Config
	Redis              *redis.Client
	RedisCtx           context.Context
	RelationalDatabase *gorm.DB

	Validate   *validator.Validate
	Translator ut.Translator
	BreakError constant.BreakError
	//Hbase
)
