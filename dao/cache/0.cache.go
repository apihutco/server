package cache

import (
	"fmt"

	"github.com/apihutco/server/config"
	"github.com/go-redis/redis"
)

type Cache struct {
	redis *redis.Client
}

var sharedCache *Cache

func Init() (err error) {
	sharedCache = new(Cache)
	cfg := config.Conf.Redis
	if !cfg.Enable {
		return nil
	}
	sharedCache.redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	_, err = sharedCache.redis.Ping().Result()
	return err
}

func Ctrl() *Cache {
	if sharedCache == nil {
		if err := Init(); err != nil {
			panic(err)
		}
	}
	return sharedCache
}
