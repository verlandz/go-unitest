package simple

// Calc just sample simple function to show looping test in unitest.
func Calc(a, b int) int {
	if a < 0 {
		a *= -1
	}
	return a + b
}
