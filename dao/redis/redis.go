package redis

import (
	"bluebell/settings"
	"fmt"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

var RDB *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	RDB = redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf("%s:%d",
				cfg.Host,
				cfg.Port,
			),
			Password: cfg.Password,
			DB:       cfg.DBName,
			PoolSize: cfg.PoolSize, //
		})
	_, err = RDB.Ping().Result()
	if err != nil {
		zap.L().Error("Connect Redis Failed", zap.Error(err))
	}
	return nil
}
func Close() {
	_ = RDB.Close()
	zap.L().Info("Redis Close")
}
