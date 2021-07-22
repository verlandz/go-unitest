package component

import (
	"testing"

	"github.com/golang/mock/gomock"
	tt "github.com/verlandz/go-unitest/pkg/tester"
	tCalc "github.com/verlandz/go-unitest/tests/calc"
	rCalcRedis "github.com/verlandz/go-unitest/unitest-sample/dependency-injection/repository/calc/redis"
	rCalcRedisMocks "github.com/verlandz/go-unitest/unitest-sample/dependency-injection/repository/calc/redis/mocks"
)

type Test struct{}

// TestDoCalcFailWhenGetNumberIsFail fail when calcRedis.GetNumber is fail.
func (Test) TestDoCalcFailWhenGetNumberIsFail(t *testing.T) (actual int, err error) {
	mockN := tCalc.GetN()
	mockCalcRedis := rCalcRedisMocks.NewMockClient(gomock.NewController(t))

	t.Run(tt.Name{
		Given: "number",
		When:  "calcRedis.GetNumber is fail",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		mockCalcRedis.EXPECT().
			GetNumber(120).
			Return(rCalcRedis.Test{}.TestGetNumberFailWhenRedisIsNil(t))

		u := New(mockCalcRedis)
		actual, err = u.DoCalc(mockN)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	return actual, err
}

// TestDoCalcSuccess success when everything is ok.
func (Test) TestDoCalcSuccess(t *testing.T) (actual int, err error) {
	mockN := tCalc.GetN()
	mockCalcRedis := rCalcRedisMocks.NewMockClient(gomock.NewController(t))

	t.Run(tt.Name{
		Given: "number",
		When:  "calcRedis.GetNumber is ok",
		Then:  "return data and no err",
	}.Construct(), func(t *testing.T) {
		mockCalcRedis.EXPECT().
			GetNumber(120).
			Return(rCalcRedis.Test{}.TestGetNumberSuccess(t))

		u := New(mockCalcRedis)
		actual, err = u.DoCalc(mockN)
		expected := 130

		tt.Equal(t, expected, actual)
		tt.NoError(t, err)
	})

	return actual, err
}
