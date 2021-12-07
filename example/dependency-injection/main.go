package main

import (
	"log"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/paytm/grace.v1"

	// delivery
	dNumHttp "github.com/verlandz/go-unitest/example/dependency-injection/delivery/num/http"
	// usecase
	uNumComponent "github.com/verlandz/go-unitest/example/dependency-injection/usecase/num/component"
	// repository
	rNumHttp "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/http"
)

func main() {
	// HTTP Router
	router := httprouter.New()

	// Repository
	numHttpClient := rNumHttp.New(&http.Client{})
	// Usecase
	numComponentClient := uNumComponent.New(numHttpClient)
	// Delivery
	_ = dNumHttp.New(router, numComponentClient)

	// Listen
	log.Fatal(grace.Serve(":8080", router))
}
