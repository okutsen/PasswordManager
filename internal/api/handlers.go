package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/okutsen/PasswordManager/internal/log"
)

const (
	RecordCreatedMessage  = "Record created"
)

type HandlerContext struct {
	ctrl   Controller
	logger log.Logger
}

func NewGetAllRecordsHandler(hctx *HandlerContext) httprouter.Handle {
	contextLogger := hctx.logger.WithFields(log.Fields{"handler": "getAllRecords"})
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		responseBody, err := hctx.ctrl.GetAllRecords()
		if err != nil {
			contextLogger.Warnf("failed to get responce body: %s", err.Error())
			writeResponse(responseBody, http.StatusInternalServerError, w, contextLogger)
			return
		}
		writeResponse(responseBody, http.StatusOK, w, contextLogger)
	}
}

func NewGetRecordHandler(hctx *HandlerContext) httprouter.Handle {
	contextLogger := hctx.logger.WithFields(log.Fields{"handler": "getRecord"})
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		idStr := ps.ByName(IDParamName)
		idInt, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			contextLogger.Warnf("failed to convert path parameter id: %s", err.Error())
			code := http.StatusBadRequest
			writeResponse(http.StatusText(code), code, w, contextLogger)
			return
		}
		responseBody, err := hctx.ctrl.GetRecord(idInt)
		if err != nil {
			contextLogger.Warnf("failed to get responce body: %s", err.Error())
			code := http.StatusInternalServerError
			writeResponse(http.StatusText(code), code, w, contextLogger)
			return
		}
		writeResponse(responseBody, http.StatusOK, w, contextLogger)
	}
}

func NewCreateRecordsHandler(hctx *HandlerContext) httprouter.Handle {
	contextLogger := hctx.logger.WithFields(log.Fields{"handler": "createRecords"})
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		responseBody, err := hctx.ctrl.GetAllRecords()
		if err != nil {
			contextLogger.Warnf("failed to get responce body: %s", err.Error())
			writeResponse(responseBody, http.StatusAccepted, w, contextLogger)
			return
		}
		writeResponse(RecordCreatedMessage, http.StatusAccepted, w, contextLogger)
	}
}

func writeResponse(body string, statusCode int, w http.ResponseWriter, logger log.Logger) {
	// TODO: write JSON response
	w.WriteHeader(statusCode)
	_, err := fmt.Fprint(w, body)
	if err != nil {
		logger.Warnf("failed to write response: %s", err.Error())
	}
	logger.Infof("response written\n%s", body)
}
