package internal

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

type Config struct {
	RedisAddress  string
	RedisPassword string
}

func New(config Config) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword,
		DB:       0,
	})

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	return &RedisClient{
		client: client,
	}, nil
}

func (r *RedisClient) WriteString(key, value string, expiration time.Duration) error {
	if err := r.client.Set(key, value, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) ReadString(key string) (string, error) {
	value, err := r.client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}
