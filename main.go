package main

import (
	"log"

	"github.com/julienschmidt/httprouter"
	pkgRedis "github.com/verlandz/go-unitest/pkg/redis"
	dCalcHttp "github.com/verlandz/go-unitest/unitest-sample/dependency-injection/delivery/calc/http"
	rCalcRedis "github.com/verlandz/go-unitest/unitest-sample/dependency-injection/repository/calc/redis"
	uCalcComponent "github.com/verlandz/go-unitest/unitest-sample/dependency-injection/usecase/calc/component"
	"gopkg.in/paytm/grace.v1"
)

func main() {
	// Redis
	myredis, err := pkgRedis.New("127.0.0.1:6379", "")
	if err != nil {
		log.Println(err)
	}

	// Http Router
	router := httprouter.New()

	// Repository
	calcRedis := rCalcRedis.New(myredis)
	// Usecase
	calcComponent := uCalcComponent.New(calcRedis)
	// Delivery
	_ = dCalcHttp.New(router, calcComponent)

	// Http Graceful Serve
	log.Fatal(grace.Serve(":8080", router))
}
