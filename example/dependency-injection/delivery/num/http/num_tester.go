package http

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	tdNum "github.com/verlandz/go-unitest/example/dependency-injection/test-data/num"
	uNumComponent "github.com/verlandz/go-unitest/example/dependency-injection/usecase/num/component"
	uNumComponentMocks "github.com/verlandz/go-unitest/example/dependency-injection/usecase/num/component/mocks"
)

type TestGetLuckyNumber struct{}

// FailCalcLuckyNumber fail when CalcLuckyNumber.
func (TestGetLuckyNumber) FailCalcLuckyNumber(t *testing.T) *httptest.ResponseRecorder {
	mockCalcComponent := uNumComponentMocks.NewMockClient(gomock.NewController(t))
	mockN := tdNum.GetNumber()

	mockCalcComponent.EXPECT().
		CalcLuckyNumber(mockN).
		Return(uNumComponent.TestCalcLuckyNumber{}.FailGetTodayNumber(t))

	router := httprouter.New()
	request, _ := http.NewRequest(http.MethodGet, "/v1/lucky-number", nil)
	request.Form = make(url.Values)
	request.Form.Add("n", strconv.Itoa(mockN))
	response := httptest.NewRecorder()

	New(router, mockCalcComponent)
	router.ServeHTTP(response, request)
	return response
}

// Success when everything is success.
func (TestGetLuckyNumber) Success(t *testing.T) *httptest.ResponseRecorder {
	mockCalcComponent := uNumComponentMocks.NewMockClient(gomock.NewController(t))
	mockN := tdNum.GetNumber()

	mockCalcComponent.EXPECT().
		CalcLuckyNumber(mockN).
		Return(uNumComponent.TestCalcLuckyNumber{}.Success(t))

	router := httprouter.New()
	request, _ := http.NewRequest(http.MethodGet, "/v1/lucky-number", nil)
	request.Form = make(url.Values)
	request.Form.Add("n", strconv.Itoa(mockN))
	response := httptest.NewRecorder()

	New(router, mockCalcComponent)
	router.ServeHTTP(response, request)
	return response
}
