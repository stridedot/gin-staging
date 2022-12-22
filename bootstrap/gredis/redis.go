package gredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go_code/gintest/bootstrap/glog"
	"strconv"
	"time"
)

var Rdb *redis.Client

type RedisConfig struct {
	Host        string `mapstructure:"host"`
	Password    string `mapstructure:"password"`
	Port        int    `mapstructure:"port"`
	Db          int    `mapstructure:"db"`
	PoolSize    int    `mapstructure:"pool_size"`
	MinIdleConn int    `mapstructure:"min_idle_conn"`
}

func InitRedis(config *RedisConfig) {
	Rdb = redis.NewClient(&redis.Options{
		Addr: config.Host + ":" + strconv.Itoa(config.Port),
		Password: config.Password,
		DB: config.Db,
		PoolSize: config.PoolSize,
		MinIdleConns: config.MinIdleConn,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	var err error
	_, err = Rdb.Ping(ctx).Result()

	if err != nil {
		glog.SL.Error("Failed to create redis pool", err)
	}
}

func Close() {
	defer func(Rdb *redis.Client) {
		_ = Rdb.Close()
	}(Rdb)
}


