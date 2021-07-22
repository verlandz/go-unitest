package redis

import (
	"github.com/go-redis/redis"
)

// New creates redis client connection.
func New(addr, pass string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
	})
	_, err := client.Ping().Result()
	return client, err
}
