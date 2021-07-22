package simple

import (
	"testing"

	tt "github.com/verlandz/go-unitest/pkg/tester"
)

func TestCalc(t *testing.T) {
	type in struct {
		a, b int
	}

	var tcs = []struct {
		name     string
		input    in
		expected int
	}{
		{
			name: tt.Name{
				Given: "a=1 and b=2",
				When:  "a is positive number",
				Then:  "return 3",
			}.Construct(),
			input: in{
				a: 1,
				b: 2,
			},
			expected: 3,
		},
		{
			name: tt.Name{
				Given: "a=-1 and b=2",
				When:  "a is negative number",
				Then:  "return 3",
			}.Construct(),
			input: in{
				a: -1,
				b: 2,
			},
			expected: 3,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := Calc(tc.input.a, tc.input.b)
			tt.Equal(t, tc.expected, actual)
		})
	}
}
