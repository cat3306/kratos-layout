package thirdmodule

import (
	"context"
	"time"

	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/redis/go-redis/v9"
)

func initRedis(config *conf.ThirdModule_Redis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         config.Addr,
		Network:      config.Network,
		Username:     "",
		Password:     config.Password,
		DB:           int(config.Db),
		DialTimeout:  time.Duration(config.DialTimeoutSeconds) * time.Second,
		ReadTimeout:  config.ReadTimeout.AsDuration(),
		WriteTimeout: config.WriteTimeout.AsDuration(),
		PoolSize:     int(config.PoolSize),
	})
	if config.PingTimeoutSeconds == 0 {
		config.PingTimeoutSeconds = 2
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.PingTimeoutSeconds)*time.Second)
	defer cancel()
	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return client, nil
}
