package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

var Client *redis.Client

const Nil = redis.Nil

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})

	fmt.Println("redis addr:", viper.GetString("redis.addr"))

	_, err := Client.Ping().Result()
	if err != nil {
		log.Errorf(err, "[redis] redis ping err")
		panic(err)
	}
}