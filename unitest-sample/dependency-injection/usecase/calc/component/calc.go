package component

import (
	"log"
)

// DoCalc do some calculation from given input.
func (u *usecase) DoCalc(n int) (int, error) {
	n *= 10
	n += 20

	res, err := u.calcRedis.GetNumber(n)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	n += res
	return n, nil
}
