package redis

import (
	"encoding/json"
	"errors"
	"fmt"
)

// REDIS_GET_NUMBER_KEY is redis key to get the number.
//
// assuming your r.FormValue("n")=0, try do this cmd on redis:
//	set get:number:20 10
const REDIS_GET_NUMBER_KEY = "get:number:%v"

// GetNumber get a number from redis base on given n.
func (r *repository) GetNumber(n int) (int, error) {
	res := 0
	if r.redis == nil {
		return res, errors.New("redis is nil")
	}

	key := fmt.Sprintf(REDIS_GET_NUMBER_KEY, n)
	temp, err := r.redis.Get(key).Result()
	if err != nil {
		return res, err
	}

	json.Unmarshal([]byte(temp), &res)
	return n, nil
}
