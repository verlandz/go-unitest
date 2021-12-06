package main

import (
	"testing"

	tt "github.com/verlandz/go-pkg/tester"
)

func TestGetSquareNumber(t *testing.T) {
	type out struct {
		value int
		isErr bool
	}

	var tcs = []struct {
		name     string
		input    int
		expected out
	}{
		{
			name: tt.Name{
				Given: "number",
				When:  "it's positive number",
				Then:  "return new value and no err",
			}.Construct(),
			input: 2,
			expected: out{
				value: 4,
				isErr: false,
			},
		},
		{
			name: tt.Name{
				Given: "number",
				When:  "it's negative number",
				Then:  "return zero value and err",
			}.Construct(),
			input: -2,
			expected: out{
				value: 0,
				isErr: true,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GetSquareNumber(tc.input)
			tt.Equal(t, tc.expected.value, actual)

			if tc.expected.isErr {
				tt.Error(t, err)
			} else {
				tt.NoError(t, err)
			}
		})
	}
}
