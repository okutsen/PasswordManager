package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

const (
	testSeparator string = "\n---------------------\n"
)

type TableTest struct {
	name               string
	httpPathParam      string
	expectedHTTPStatus int
	expectedBody       string
}

type TableTests struct {
	tt         []*TableTest // by pointer or not?
	httpMethod string
	httpPath   string
	httpServer *ClientAPI
	handler    HandlerFunc
}

type HandlerFunc func(*ClientAPI, http.ResponseWriter, *http.Request, httprouter.Params)

func TestGetRecords(t *testing.T) {
	tests := TableTests{
		tt: []*TableTest{
			{
				name:               "Get all records",
				httpPathParam:      "",
				expectedHTTPStatus: http.StatusOK,
				expectedBody:       "Records:\n0,1,2,3,4,5",
			},
			{
				name:               "Get record by Name:1",
				httpPathParam:      "0",
				expectedHTTPStatus: http.StatusOK,
				expectedBody:       "Records:\n0",
			},
			{
				name:               "Get record by Name:5",
				httpPathParam:      "5",
				expectedHTTPStatus: http.StatusOK,
				expectedBody:       "Records:\n5",
			},
			{
				name:               "Returns 404 on missing record",
				httpPathParam:      "6",
				expectedHTTPStatus: http.StatusBadRequest,
				expectedBody:       "Record not found",
			}},

		httpMethod: http.MethodPost,
		httpPath:   "/records/",
		httpServer: NewClientAPI(),
		handler:    (*ClientAPI).getRecords,
	}
	TableTestRunner(t, tests)
}

func TestPostRecords(t *testing.T) {
	tests := TableTests{
		tt: []*TableTest{
			{
				name:               "Post record",
				expectedHTTPStatus: http.StatusAccepted,
				expectedBody:       "New Record created",
			}},

		httpMethod: http.MethodPost,
		httpPath:   "/records/",
		httpServer: NewClientAPI(),
		handler:    (*ClientAPI).createRecords,
	}
	TableTestRunner(t, tests)
}

func TableTestRunner(t *testing.T, tt TableTests) {
	t.Helper()
	for _, test := range tt.tt {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.httpMethod, tt.httpPath+test.httpPathParam, nil)
			response := httptest.NewRecorder()
			ps := httprouter.Params{httprouter.Param{Key: getByIdParamName, Value: test.httpPathParam}}
			tt.handler(tt.httpServer, response, request, ps)

			assert(t, response.Code, test.expectedHTTPStatus, "Wrong status")
			assert(t, response.Body.String(), test.expectedBody, "Wrong body")
		})
	}
}

func assert[T comparable](t *testing.T, got, want T, errorMessage string) {
	t.Helper()
	if got != want {
		t.Errorf("%s\nGot:%s%v%sWant:%s%v%s",
			errorMessage,
			testSeparator, got, testSeparator,
			testSeparator, want, testSeparator)
	}
}
