package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/rs/zerolog/log"
)

func Redis() {
	opt := redis.Options{
		Addr:     global.Configuration.Redis.TcpAddr.String(),
		Username: global.Configuration.Redis.User,
		Password: global.Configuration.Redis.Password,
		DB:       global.Configuration.Redis.DB,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			log.Info().Msg("Connected to Redis")
			return nil
		},
	}
	global.RedisCtx = context.Background()
	global.Redis = redis.NewClient(&opt)
}
