package mocks

import (
	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
)

// NewMock creates mock client connection.
func NewMock() *redismock.ClientMock {
	mr, _ := miniredis.Run()
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return redismock.NewNiceMock(client)
}
