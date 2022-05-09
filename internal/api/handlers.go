package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/okutsen/PasswordManager/domain"
	"github.com/okutsen/PasswordManager/internal/log"
)

const (
	RecordCreatedMessage = "Record created"
)

func NewEndpointLoggerMiddleware(ctx *APIContext, handler httprouter.Handle) httprouter.Handle {
		return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			ctx.logger.Infof("API: Endpoint Hit: %s %s%s\n", r.Host, r.URL.Path, r.Method)
			handler(rw, r, ps)
		}
	}

func NewGetAllRecordsHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "getAllRecords"})
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		records, err := ctx.ctrl.GetAllRecords()
		if err != nil {
			logger.Warnf("failed to get response from controller: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		writeJSONResponse(w, logger, records)
	}
}

func NewGetRecordHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "getRecord"})
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		idStr := ps.ByName(IDParamName)
		idInt, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			logger.Warnf("failed to convert path parameter id: %s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		records, err := ctx.ctrl.GetRecord(idInt)
		if err != nil {
			logger.Warnf("failed to get response from controller: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		writeJSONResponse(w, logger, records)
	}
}

func NewCreateRecordsHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "createRecords"})
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// TODO: check content type
		records, err := readRecords(r.Body)
		defer r.Body.Close()
		if err != nil {
			logger.Warnf("failed to parse JSON: %s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = ctx.ctrl.CreateRecords(records)
		if err != nil {
			logger.Warnf("failed to get response from controller: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		writeTextResponse(w, logger, RecordCreatedMessage, http.StatusAccepted)
	}
}

func readRecords(requestBody io.Reader) ([]domain.Record, error) {
	// TODO: prevent overflow (read by batches or set max size)
	recordsJSON, err := ioutil.ReadAll(requestBody)
	if err != nil {
		return nil, err
	}
	var records []domain.Record
	err = json.Unmarshal(recordsJSON, &records)
	if err != nil {
		return nil, err
	}
	return records, err
}

func writeTextResponse(w http.ResponseWriter, logger log.Logger, body string, statusCode int) {
	_, err := fmt.Fprint(w, body)
	if err != nil {
		logger.Warnf("failed to write text response: %s", err.Error())
	}
	w.WriteHeader(statusCode)
	logger.Infof("response written\n%s", body)
}

func writeJSONResponse(w http.ResponseWriter, logger log.Logger, body []domain.Record) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		logger.Warnf("failed to write JSON response: %s", err.Error())
	}
	// TODO: do not log private info
	logger.Infof("response written: %+v", body)
}
