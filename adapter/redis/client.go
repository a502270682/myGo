package redis

import (
	"github.com/go-redis/redis/v8"
	"sync"
)

const RedisClusterName = "default"

var client *RedisClient

var ErrNil = redis.Nil

type RedisClient struct {
	sync.RWMutex
	clients map[string]*redis.Client
}

func (c *RedisClient) redisClientByName(name string) *redis.Client {
	c.RLock()
	cli, ok := c.clients[name]
	c.RUnlock()
	if !ok {
		return nil
	}
	return cli
}

func GetDefaultRedisClient() *redis.Client {
	return client.redisClientByName(RedisClusterName)
}
