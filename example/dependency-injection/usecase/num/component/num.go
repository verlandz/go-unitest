package component

import (
	"log"
)

// CalcLuckyNumber calculates lucky number from given input.
func (u *usecase) CalcLuckyNumber(n int) (int, error) {
	n = (n * 10) + 10 // random logic (1)

	num, err := u.numRedis.GetTodayNumber(n)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	n += num // random logic (2)
	return n, nil
}
