package db

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

var Client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	Password: "",
	DB: 0,
})

var (
	client = &redisClient{}
)

type redisClient struct {
	c *redis.Client
}

//GetClient get the redis client
func Initialize() *redisClient {
	c := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	if err := c.Ping(c.Context()).Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}
	client.c = c
	return client
}

//SetKey set key
func (client *redisClient) SetKey(c *gin.Context, key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return err
	}
	//err = client.c.Set(key, cacheEntry, expiration).Err()
	err = client.c.Set(c, key, cacheEntry, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (client *redisClient) DelKey(c *gin.Context,key string) error {
	return client.c.Del(c,key).Err()
}