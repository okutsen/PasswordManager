package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/okutsen/PasswordManager/internal/log"
)

func NewGetAllRecordsHandler(logger log.Logger) httprouter.Handle {
	contextLogger := logger.WithFields(log.Fields{"handler": "getAllRecords"})
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		// TODO: get from controller
		responseBody := "Records:\n0,1,2,3,4,5"
		writeResponse(responseBody, http.StatusOK, w, contextLogger)
	}
}

func NewGetRecordHandler(logger log.Logger) httprouter.Handle {
	contextLogger := logger.WithFields(log.Fields{"handler": "getRecord"})
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		idStr := ps.ByName(IDParamName)
		// TODO: convert to correct type (uint)
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			contextLogger.Warnf("Failed to convert path parameter id: %s", err.Error())
			writeResponse(RecordNotFoundMessage, http.StatusBadRequest, w, contextLogger)
			return
		}
		responseBody := fmt.Sprintf("Records:\n%d", idInt)
		writeResponse(responseBody, http.StatusOK, w, contextLogger)
	}
}



func NewCreateRecordsHandler(logger log.Logger) httprouter.Handle {
	contextLogger := logger.WithFields(log.Fields{"handler": "createRecords"})
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		writeResponse(RecordCreatedMessage, http.StatusAccepted, w, contextLogger)
	}
}

func writeResponse(body string, status int, w http.ResponseWriter, logger log.Logger) error {
	// TODO: write JSON response
	w.WriteHeader(status)
	_, err := fmt.Fprint(w, body)
	if err != nil {
		return err
	}
	logger.Infof("Response written\n%s", body)
	return nil
}
