package internal

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestGetRecords(t *testing.T) {
	client := NewClientAPI()
	tests := []struct {
		name               string
		recordName         string
		expectedHTTPStatus int
		expectedResult     string
	}{
		{
			name:               "Get record by Name:1",
			recordName:         "",
			expectedHTTPStatus: http.StatusOK,
			expectedResult:     "Records:\n0,1,2,3,4,5",
		},
		{
			name:               "Get record by Name:1",
			recordName:         "0",
			expectedHTTPStatus: http.StatusOK,
			expectedResult:     "Records:\n0",
		},
		{
			name:               "Get record by Name:5",
			recordName:         "5",
			expectedHTTPStatus: http.StatusOK,
			expectedResult:     "Records:\n5",
		},
		{
			name:               "Returns 404 on missing record",
			recordName:         "6",
			expectedHTTPStatus: http.StatusBadRequest,
			expectedResult:     "Record not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := newGetRequest(tt.recordName)
			response := httptest.NewRecorder()
			ps := httprouter.Param{Key: "recordName", Value: tt.recordName}
			client.getRecords(response, request, httprouter.Params{ps})

			assertStatus(t, response.Code, tt.expectedHTTPStatus)
			assertResponseBody(t, response.Body.String(), tt.expectedResult)
		})
	}

}

func newGetRequest(recordName string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/records/%s", recordName), nil)
	return req
}

func newPostRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/records"), nil)
	return req
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
