package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/yyyThree/zap/helper"
	"time"
)

var clients = make(map[string]*redis.Client)

type Config struct {
	Host           string
	Port           int
	User           string
	Password       string
	Db             int
	ConnectTimeout int
	ReadTimeout    int
	WriteTimeout   int
	PoolSize       int
}

func GetConn(config Config) (*redis.Client, error) {
	configMd5 := helper.StructToMd5(config)
	if client, ok := clients[configMd5]; ok && (client != nil && client.Ping(GetCtx()).Err() == nil) {
		return client, nil
	}

	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		Username:     config.User,
		Password:     config.Password,
		DB:           config.Db,
		DialTimeout:  time.Duration(config.ConnectTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
		PoolSize:     config.PoolSize,
	})
	if err := client.Ping(GetCtx()).Err(); err != nil {
		return nil, err
	}
	clients[configMd5] = client
	return client, nil
}

func GetCtx() context.Context {
	return context.Background()
}
