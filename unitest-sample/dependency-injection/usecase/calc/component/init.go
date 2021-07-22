package component

import (
	rCalcRedis "github.com/verlandz/go-unitest/unitest-sample/dependency-injection/repository/calc/redis"
)

type Client interface {
	DoCalc(n int) (int, error)
}

type usecase struct {
	calcRedis rCalcRedis.Client
}

// New creates usecase Client.
func New(calcRedis rCalcRedis.Client) Client {
	return &usecase{
		calcRedis: calcRedis,
	}
}
