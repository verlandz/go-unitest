package main

import (
	"log"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/paytm/grace.v1"

	// delivery
	dNumHttp "github.com/verlandz/go-unitest/example/dependency-injection/delivery/num/http"
	// usecase
	uNumComponent "github.com/verlandz/go-unitest/example/dependency-injection/usecase/num/component"
	// repository
	rNumGenerate "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/generate"
)

func main() {
	// HTTP Router
	router := httprouter.New()

	// Repository
	numGenerate := rNumGenerate.New()
	// Usecase
	numComponent := uNumComponent.New(numGenerate)
	// Delivery
	_ = dNumHttp.New(router, numComponent)

	// Http Graceful Serve
	log.Fatal(grace.Serve(":8080", router))
}
