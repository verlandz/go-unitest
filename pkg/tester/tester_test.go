package tester

import (
	"testing"
)

func TestConstruct(t *testing.T) {
	t.Run(Name{
		Given: "param",
		When:  "param is complete",
		Then:  "return complete string",
	}.Construct(), func(t *testing.T) {
		tc := Name{
			Given: "a",
			When:  "b",
			Then:  "c",
		}

		actual := tc.Construct()
		expected := "Given a. When b. Then c."

		Equal(t, expected, actual)
	})

	t.Run(Name{
		Given: "param",
		When:  "param is not complete",
		Then:  "return complete string",
	}.Construct(), func(t *testing.T) {
		tc := Name{
			Given: "a",
		}

		Panics(t, func() {
			tc.Construct()
		})
	})
}
