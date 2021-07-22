package redis

import (
	"testing"

	"github.com/alicebob/miniredis"
	tt "github.com/verlandz/go-unitest/pkg/tester"
)

func TestNew(t *testing.T) {
	t.Run(tt.Name{
		Given: "addr and pass",
		When:  "both is empty",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		_, err := New("", "")
		tt.Error(t, err)
	})

	t.Run(tt.Name{
		Given: "addr and pass",
		When:  "addr is exist",
		Then:  "return data and no err",
	}.Construct(), func(t *testing.T) {
		mr, _ := miniredis.Run()
		addr := mr.Addr()

		_, err := New(addr, "")
		tt.NoError(t, err)
	})
}
