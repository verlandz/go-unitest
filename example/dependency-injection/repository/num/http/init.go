package http

import "net/http"

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/Client.go -package=mocks . Client

type Client interface {
	GetRandNumber(n int) (int, error)
}

type repository struct {
	client *http.Client
}

func New(client *http.Client) Client {
	return &repository{
		client: client,
	}
}
