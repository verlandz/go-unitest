package num

import (
	"testing"

	tt "github.com/verlandz/go-pkg/tester"
)

func TestGetPositiveN(t *testing.T) {
	var tcs = []struct {
		name     string
		expected int
	}{
		{
			name: tt.Name{
				Given: "nothing",
				When:  "data is present",
				Then:  "return positive number",
			}.Construct(),
			expected: 10,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := GetPositiveN()
			tt.Equal(t, tc.expected, actual)
		})
	}
}

func TestGetNegativeN(t *testing.T) {
	var tcs = []struct {
		name     string
		expected int
	}{
		{
			name: tt.Name{
				Given: "nothing",
				When:  "data is present",
				Then:  "return negative number",
			}.Construct(),
			expected: -10,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := GetNegativeN()
			tt.Equal(t, tc.expected, actual)
		})
	}
}
