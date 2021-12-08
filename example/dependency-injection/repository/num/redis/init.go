package redis

import (
	"github.com/go-redis/redis"
)

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/Client.go -package=mocks . Client

type Client interface {
	GetTodayNumber(n int) (int, error)
}

type repository struct {
	client redis.Cmdable
}

func New(client redis.Cmdable) Client {
	return &repository{
		client: client,
	}
}
