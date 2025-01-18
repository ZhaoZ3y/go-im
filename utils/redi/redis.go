package redi

import (
	"context"
	"github.com/redis/go-redis/v9"
	"goim/config"
	"strconv"
	"time"
)

var RDB *redis.Client

func RedisInit() {
	addr := config.GetConfig().Redis.Host + ":" + strconv.Itoa(config.GetConfig().Redis.Port)
	password := config.GetConfig().Redis.Password
	db := config.GetConfig().Redis.DB

	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	RDB = redisClient
}

func Get(key string) (val string) {
	return RDB.Get(context.Background(), key).Val()
}

func Set(key string, value string) error {
	return RDB.Set(context.Background(), key, value, 5*time.Minute).Err()
}
