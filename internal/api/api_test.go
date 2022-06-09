package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/okutsen/PasswordManager/config"
	"github.com/okutsen/PasswordManager/internal/controller"
	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/apischema"
	"github.com/stretchr/testify/assert"
)

const (
	testSeparator string = "\n---------------------\n"
)

func initAPI() {
	var logger log.Logger = log.NewLogrusLogger()
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatalf("Failed to initialize config: %v", err)
	}
	ctrl := controller.New(logger)
	serviceAPI := New(&Config{Port: cfg.Port}, ctrl, logger)

	go func() {
		err = serviceAPI.Start()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("Failed to start application %v", err)
			return
		}
	}()
}

type TableTest struct {
	// Create request for specified
	testName           string
	httpMethod         string
	httpPath           string
	expectedHTTPStatus int
	expectedBody       interface{}
}

type TableTests struct {
	tt []*TableTest
}

func TestGetRecords(t *testing.T) {
	tests := TableTests{
		tt: []*TableTest{
			{
				testName:           "Get all records",
				httpMethod:         http.MethodGet,
				httpPath:           "/records",
				expectedHTTPStatus: http.StatusOK,
				// TODO: fill with test data
				expectedBody:       &apischema.Record{

				},
			},
			{
				testName:   "Get record by id 1",
				httpMethod: http.MethodGet,
				httpPath:   "/records/0",
				expectedHTTPStatus: http.StatusOK,
				// TODO: fill with test data
				expectedBody:       &apischema.Record{

				},
			},
			{
				testName:   "Get record by id 5",
				httpMethod: http.MethodGet,
				httpPath:   "/records/5",
				expectedHTTPStatus: http.StatusOK,
				// TODO: fill with test data
				expectedBody:       &apischema.Record{

				},
			},
			{
				testName:   "Returns 404 on missing record",
				httpMethod: http.MethodGet,
				httpPath:   "/records/a",
				expectedHTTPStatus: http.StatusBadRequest,
				// TODO: fill with test data
				expectedBody:       &apischema.Error{

				},
			}},
	}
	TableTestRunner(t, tests)
}

func TestPostRecords(t *testing.T) {
	tests := TableTests{
		tt: []*TableTest{
			{
				testName:           "Post record",
				httpMethod:         http.MethodPost,
				httpPath:           "/records/",
				expectedHTTPStatus: http.StatusAccepted,
				expectedBody:       "", // workaround
			}},
	}
	TableTestRunner(t, tests)
}

func TableTestRunner(t *testing.T, tt TableTests) {
	t.Helper()
	initAPI()
	// Use generated client
	cl := http.Client{
		Timeout: 6 * time.Second,
	}
	assert := assert.New(t)
	for _, test := range tt.tt {
		t.Run(test.testName, func(t *testing.T) {

			request := httptest.NewRequest(test.httpMethod, test.httpPath, nil)
			// response := httptest.NewRecorder()
			response, err := cl.Do(request)
			assert.Nil(err, "Response should be received")
			
			var receivedBody apischema.Record
			err = readJSON(response.Body, &receivedBody)
			assert.Nil(err, "Response body should match object schema")

			assert.Equal(response.StatusCode, test.expectedHTTPStatus, "Wrong status")
			assert.Equal(receivedBody, test.expectedBody, "Wrong body")
		})
	}
}
