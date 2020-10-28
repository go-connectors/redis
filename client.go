package redis

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

// Client is a Redis client representing a pool of zero or more
// underlying connections.
type Client struct {
	*redis.Client
	Config Config
	ctx    context.Context
}

// NewClient creates connection to Redis and returns Redis Client.
func NewClient(cfg *Config) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		MaxRetries:   cfg.MaxRetries,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		PoolSize:     cfg.PoolSize,
	})

	ctx := context.Background()

	if _, err := client.Ping(ctx).Result(); err != nil {
		_ = client.Close()

		return nil, err
	}

	return &Client{Client: client, Config: *cfg, ctx: ctx}, nil
}

// IsErrorNotFound checks error for Redis Nil.
func IsErrorNotFound(err error) bool {
	return errors.Is(err, redis.Nil)
}
