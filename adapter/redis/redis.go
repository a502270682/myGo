package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"myGo/adapter/log"
	"time"
)

type RedisConf struct {
	Name string

	Network string
	// host:port address.
	Addr               string
	Username           string
	Password           string
	DB                 int
	MaxRetries         int
	MinRetryBackoff    time.Duration
	MaxRetryBackoff    time.Duration
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	PoolSize           int
	MinIdleConns       int
	MaxConnAge         time.Duration
	PoolTimeout        time.Duration
	IdleTimeout        time.Duration
	IdleCheckFrequency time.Duration

	// Enables read only queries on slave nodes.
	readOnly bool
}

var configMap map[string]*RedisConf

func init() {
	client = &RedisClient{
		clients: make(map[string]*redis.Client),
	}
	configMap = make(map[string]*RedisConf)
}

// Initialize the redis connections based configuration
func Initialize(config *RedisConf) error {
	if config == nil {
		panic(errors.New("redis cfg is nil"))
	}
	name := config.Name
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
		Username: config.Username,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return errors.Wrapf(err, "can't conn to redis: %s", config.Addr)
	}

	if _, ok := client.clients[name]; !ok {
		client.clients[name] = rdb
		configMap[name] = config
	} else {
		log.Warnf(context.Background(), "failed to initialize redis connection to %v", name)
		return errors.New("redis is not available")
	}
	return nil
}
