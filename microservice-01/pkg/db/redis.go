package db

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func RedisConn() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_ADR"),
		Password: viper.GetString("REDIS_PASS"),
		DB:       viper.GetInt("REDIS_DB"),
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal("Could not connect to Redis: ", err)
	} else {
		fmt.Println("Connected to Redis")
	}

	return rdb
}
