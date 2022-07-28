package redis

// redis 数据库

import (
	"context"
	"demon/common"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var (
	Rdb *redis.Client
	wg  sync.WaitGroup
	ctx = context.Background()
)

func init() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		Rdb = createClient(common.Config.Redis.Host, common.Config.Mongo.Password, common.Config.Redis.Port)
	}()
	wg.Wait()
}

func createClient(host, password string, port int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Password:     password,
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     100,
		PoolTimeout:  30 * time.Second,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("redis.init err : [%v]", err)
	} else {
		log.Infof("Redis [%s:%d] 初始化完成", host, port)
	}
	return rdb
}
