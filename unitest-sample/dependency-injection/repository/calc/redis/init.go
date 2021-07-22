package redis

import (
	"github.com/go-redis/redis"
)

type Client interface {
	GetNumber(n int) (int, error)
}

type repository struct {
	redis redis.Cmdable
}

// New creates repository Client.
func New(redis redis.Cmdable) Client {
	return &repository{
		redis: redis,
	}
}
