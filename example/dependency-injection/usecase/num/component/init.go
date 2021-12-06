package component

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/Client.go -package=mocks . Client

import (
	rNumGenerate "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/generate"
)

type Client interface {
	CalcLuckyNumber(n int) (int, error)
}

type usecase struct {
	numGenerate rNumGenerate.Client
}

func New(numGenerate rNumGenerate.Client) Client {
	return &usecase{
		numGenerate: numGenerate,
	}
}
