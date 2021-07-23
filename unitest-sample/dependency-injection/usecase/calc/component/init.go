package component

import (
	rCalcRedis "github.com/verlandz/go-unitest/unitest-sample/dependency-injection/repository/calc/redis"
)

//go:generate mockgen -source=init.go -destination=mocks/Client.go -package=mocks . Client

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
