package component

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/Client.go -package=mocks . Client

import (
	rNumHttp "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/http"
)

type Client interface {
	CalcLuckyNumber(n int) (int, error)
}

type usecase struct {
	numHttp rNumHttp.Client
}

func New(numHttp rNumHttp.Client) Client {
	return &usecase{
		numHttp: numHttp,
	}
}
