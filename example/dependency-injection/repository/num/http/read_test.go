package http

import (
	"testing"
)

func Test_GetRandNumber(t *testing.T) {
	TestGetRandNumber{}.FailClientNil(t)
	TestGetRandNumber{}.FailClientNewRequest(t)
	TestGetRandNumber{}.FailClientDo(t)
	TestGetRandNumber{}.FailClientUnmarshall(t)
	TestGetRandNumber{}.FailEmptyData(t)
	TestGetRandNumber{}.Success(t)
}
