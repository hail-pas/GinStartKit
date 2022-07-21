package global

import (
	"context"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/hail-pas/GinStartKit/config"
	"gorm.io/gorm"
)

var (
	Configuration      *config.Config
	Redis              *redis.Client
	RedisCtx           context.Context
	RelationalDatabase *gorm.DB

	Validate *validator.Validate
	Trans    ut.Translator
	//Hbase
)
