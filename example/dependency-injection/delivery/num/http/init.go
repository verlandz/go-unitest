package http

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/Client.go -package=mocks . Client

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	uNumComponent "github.com/verlandz/go-unitest/example/dependency-injection/usecase/num/component"
)

type Client interface {
	GetLuckyNumber(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type delivery struct {
	numComponent uNumComponent.Client
}

func New(
	router *httprouter.Router,
	numComponent uNumComponent.Client,
) Client {
	d := &delivery{
		numComponent: numComponent,
	}

	router.GET("/v1/lucky-number", d.GetLuckyNumber)
	return d
}
