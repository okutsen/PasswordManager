package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"

	"github.com/okutsen/PasswordManager/internal/controller"
	"github.com/okutsen/PasswordManager/internal/log"
)

const (
	testSeparator string = "\n---------------------\n"
)

// FIXME: update tests for json
type TableTest struct {
	// Create request for specified
	testName   string
	handle     InnerHandlerFunc
	httpMethod string
	httpPath   string
	// TODO: constructs params with a func
	ps                 httprouter.Params
	expectedHTTPStatus int
	expectedBody       string
}

type TableTests struct {
	tt []*TableTest
}

func TestGetRecords(t *testing.T) {
	tests := TableTests{
		tt: []*TableTest{
			{
				testName:           "Get all records",
				handle:             NewGetAllRecordsHandler(),
				httpMethod:         http.MethodGet,
				httpPath:           "/records",
				expectedHTTPStatus: http.StatusOK,
				expectedBody:       "Records:\n0,1,2,3,4,5",
			},
			{
				testName:   "Get record by id 1",
				handle:     NewGetRecordHandler(),
				httpMethod: http.MethodGet,
				httpPath:   "/records/0",
				ps: httprouter.Params{
					httprouter.Param{Key: IDParamName, Value: "0"},
				},
				expectedHTTPStatus: http.StatusOK,
				expectedBody:       "Records:\n0",
			},
			{
				testName:   "Get record by id 5",
				handle:     NewGetRecordHandler(),
				httpMethod: http.MethodGet,
				httpPath:   "/records/5",
				ps: httprouter.Params{
					httprouter.Param{Key: IDParamName, Value: "5"},
				},
				expectedHTTPStatus: http.StatusOK,
				expectedBody:       "Records:\n5",
			},
			{
				testName:   "Returns 404 on missing record",
				handle:     NewGetRecordHandler(),
				httpMethod: http.MethodGet,
				httpPath:   "/records/a",
				ps: httprouter.Params{
					httprouter.Param{Key: IDParamName, Value: "a"},
				},
				expectedHTTPStatus: http.StatusBadRequest,
				expectedBody:       http.StatusText(http.StatusBadRequest),
			}},
	}
	TableTestRunner(t, tests)
}

func TestPostRecords(t *testing.T) {
	tests := TableTests{
		tt: []*TableTest{
			{
				testName:           "Post record",
				handle:             NewCreateRecordHandler(),
				httpMethod:         http.MethodPost,
				httpPath:           "/records/",
				expectedHTTPStatus: http.StatusAccepted,
				expectedBody:       "",    // workaround
			}},
	}
	TableTestRunner(t, tests)
}

func TableTestRunner(t *testing.T, tt TableTests) {
	t.Helper()
	logger := log.NewLogrusLogger()
	ctrl := controller.New(logger)
	ctx := &APIContext{ctrl, logger}
	for _, test := range tt.tt {
		t.Run(test.testName, func(t *testing.T) {
			request := httptest.NewRequest(test.httpMethod, test.httpPath, nil)
			response := httptest.NewRecorder()
			test.handle(response, request, NewRequestContext(ctx, test.ps))

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
