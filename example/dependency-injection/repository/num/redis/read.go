package redis

import (
	"errors"
	"strconv"

	"github.com/verlandz/go-unitest/example/dependency-injection/consts"
)

// GetTodayNumber returns today number.
func (r *repository) GetTodayNumber(n int) (int, error) {
	if r.client == nil {
		return 0, errors.New("client is nil")
	}

	resp, err := r.client.Get(consts.REDIS_TODAY_NUMBER_KEY).Result()
	if err != nil {
		return 0, err
	}

	newResp, err := strconv.Atoi(resp)
	if err != nil {
		return 0, err
	}

	return newResp + n, nil
}
