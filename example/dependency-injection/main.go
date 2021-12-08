package main

import (
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/julienschmidt/httprouter"
	"github.com/verlandz/go-unitest/example/dependency-injection/consts"
	"gopkg.in/paytm/grace.v1"

	// delivery
	dNumHttp "github.com/verlandz/go-unitest/example/dependency-injection/delivery/num/http"
	// usecase
	uNumComponent "github.com/verlandz/go-unitest/example/dependency-injection/usecase/num/component"
	// repository
	rNumRedis "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/redis"
)

func main() {
	// HTTP Router
	router := httprouter.New()
	// Redis Client
	redisClient := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	if _, err := redisClient.Ping().Result(); err != nil {
		panic("please start your redis")
	} else {
		// populate random data
		val := time.Now().UnixNano() % 100
		redisClient.Set(consts.REDIS_TODAY_NUMBER_KEY, val, 0)
	}

	// Repository
	numRedisClient := rNumRedis.New(redisClient)
	// Usecase
	numComponentClient := uNumComponent.New(numRedisClient)
	// Delivery
	_ = dNumHttp.New(router, numComponentClient)

	// Listen
	log.Fatal(grace.Serve(":8080", router))
}
