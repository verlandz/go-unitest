package http

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	tt "github.com/verlandz/go-unitest/pkg/tester"
	tCalc "github.com/verlandz/go-unitest/tests/calc"
	rCalcComponent "github.com/verlandz/go-unitest/unitest-sample/dependency-injection/usecase/calc/component"
	rCalcComponentMocks "github.com/verlandz/go-unitest/unitest-sample/dependency-injection/usecase/calc/component/mocks"
)

type Test struct{}

// TestCalcFailWhenDoCalcIsFail fail when calcComponent.Calc is fail.
func (Test) TestCalcFailWhenDoCalcIsFail(t *testing.T) {
	mockN := tCalc.GetN()
	mockCalcComponent := rCalcComponentMocks.NewMockClient(gomock.NewController(t))

	t.Run(tt.Name{
		Given: "n",
		When:  "calcComponent.Calc is fail",
		Then:  "return 500",
	}.Construct(), func(t *testing.T) {
		mockCalcComponent.EXPECT().
			DoCalc(100).
			Return(rCalcComponent.Test{}.TestDoCalcFailWhenGetNumberIsFail(t))

		router := httprouter.New()
		request, _ := http.NewRequest(http.MethodGet, "/v1/calc", nil)
		request.Form = make(url.Values)
		request.Form.Add("n", strconv.Itoa(mockN))
		response := httptest.NewRecorder()

		New(router, mockCalcComponent)
		router.ServeHTTP(response, request)

		expected_code := http.StatusInternalServerError
		expected_resp := `{"errors":["redis is nil"]}`

		tt.Equal(t, expected_resp, response.Body.String())
		tt.Equal(t, expected_code, response.Code)
	})
}

// TestCalcSuccess success when everything is ok.
func (Test) TestCalcSuccess(t *testing.T) {
	mockN := tCalc.GetN()
	mockCalcComponent := rCalcComponentMocks.NewMockClient(gomock.NewController(t))

	t.Run(tt.Name{
		Given: "n",
		When:  "calcComponent.Calc is fail",
		Then:  "return 200",
	}.Construct(), func(t *testing.T) {
		mockCalcComponent.EXPECT().
			DoCalc(100).
			Return(rCalcComponent.Test{}.TestDoCalcSuccess(t))

		router := httprouter.New()
		request, _ := http.NewRequest(http.MethodGet, "/v1/calc", nil)
		request.Form = make(url.Values)
		request.Form.Add("n", strconv.Itoa(mockN))
		response := httptest.NewRecorder()

		New(router, mockCalcComponent)
		router.ServeHTTP(response, request)

		expected_code := http.StatusOK
		expected_resp := `{"data":132}`

		tt.Equal(t, expected_resp, response.Body.String())
		tt.Equal(t, expected_code, response.Code)
	})
}
