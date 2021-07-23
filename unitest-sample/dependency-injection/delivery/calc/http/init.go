package http

//go:generate mockgen -source=init.go -destination=mocks/Client.go -package=mocks . Client

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	uCalcComponent "github.com/verlandz/go-unitest/unitest-sample/dependency-injection/usecase/calc/component"
)

type Client interface {
	Calc(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type delivery struct {
	calcComponent uCalcComponent.Client
}

// New creates delivery Client.
func New(
	router *httprouter.Router,
	calcComponent uCalcComponent.Client,
) Client {
	d := &delivery{
		calcComponent: calcComponent,
	}

	router.GET("/v1/calc", d.Calc)
	return d
}
