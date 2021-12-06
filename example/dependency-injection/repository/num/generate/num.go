package generate

import "errors"

// GetSquareNumber returns square of given number.
// Return error if given number is negative number.
func (r *repository) GetSquareNumber(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("n should be positive number")
	}
	return n * n, nil
}
