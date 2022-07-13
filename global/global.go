package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/hail-pas/GinStartKit/config"
	"gorm.io/gorm"
)

var (
	Configuration      *config.Config
	Redis              *redis.Client
	RedisCtx           context.Context
	RelationalDatabase *gorm.DB
	//ConcurrencyControl = &singleflight.Group{}
)
