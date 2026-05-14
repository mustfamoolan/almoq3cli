package bootstrap

import (
	"context"
	"fmt"
	"time"

	"almoq3-final-release/config"
	"github.com/fatih/color"
	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client
var Ctx = context.Background()

func InitializeCache() {
	cacheConfig := config.Global.Cache

	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cacheConfig.Host, cacheConfig.Port),
		Password: cacheConfig.Password,
		DB:       0, // use default DB
	})

	// Test connection
	_, err := Redis.Ping(Ctx).Result()
	if err != nil {
		fmt.Printf("%s Redis connection failed: %v\n", color.YellowString("⚠️"), err)
		return
	}

	fmt.Printf("%s Redis connected successfully\n", color.CyanString("⚙️"))
}

// Cache Wrapper (Laravel-style)
func CacheGet(key string) (string, error) {
	return Redis.Get(Ctx, key).Result()
}

func CacheSet(key string, value interface{}, expiration time.Duration) error {
	return Redis.Set(Ctx, key, value, expiration).Err()
}

func CacheForget(key string) error {
	return Redis.Del(Ctx, key).Err()
}
