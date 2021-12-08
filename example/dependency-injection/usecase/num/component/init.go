package component

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/Client.go -package=mocks . Client

import (
	rNumRedis "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/redis"
)

type Client interface {
	CalcLuckyNumber(n int) (int, error)
}

type usecase struct {
	numRedis rNumRedis.Client
}

func New(numRedis rNumRedis.Client) Client {
	return &usecase{
		numRedis: numRedis,
	}
}
