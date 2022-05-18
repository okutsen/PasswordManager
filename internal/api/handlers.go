package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/apischema"
	"github.com/okutsen/PasswordManager/schema/schemabuilder"
)

const (
	RecordCreatedMessage = "Record created"
)

func NewEndpointLoggerMiddleware(ctx *APIContext, handler httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx.logger.Infof("API: Endpoint Hit: %s %s%s", r.Method, r.Host, r.URL.Path)
		handler(rw, r, ps)
	}
}

func NewGetAllRecordsHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "GetAllRecords"})
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		records, err := ctx.ctrl.GetAllRecords()
		if err != nil {
			logger.Warnf("Failed to get records from controller: %s", err.Error())
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}
		recordsAPI := schemabuilder.BuildRecordsAPIFrom(records)
		writeJSONResponse(w, logger, recordsAPI, http.StatusOK)
	}
}

func NewGetRecordHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "GetRecord"})
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		idStr := ps.ByName(IDParamName)
		idInt, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			logger.Warnf("Failed to convert path parameter id: %s", err.Error())
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}
		record, err := ctx.ctrl.GetRecord(idInt)
		if err != nil {
			logger.Warnf("Failed to get records from controller: %s", err.Error())
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}
		// TODO: get record from db
		writeJSONResponse(w, logger, schemabuilder.BuildRecordAPIFrom(record), http.StatusOK)
	}
}

func NewCreateRecordHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "CreateRecords"})
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// TODO: check content type
		var recordAPI *apischema.Record
		err := readJSON(r.Body, recordAPI)
		defer r.Body.Close()
		if err != nil {
			logger.Warnf("Failed to read JSON: %s", err.Error())
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}
		record := schemabuilder.BuildRecordFrom(recordAPI)
		// TODO: if exists return err (409 Conflict)
		err = ctx.ctrl.CreateRecord(record)
		if err != nil {
			logger.Warnf("Failed to get records from controller: %s", err.Error())
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}
		// TODO: get record from db
		writeJSONResponse(w, logger, schemabuilder.BuildRecordAPIFrom(record), http.StatusAccepted)
	}
}

func NewUpdateRecordHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "UpdateRecords"})
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// TODO: check content type
		var recordAPI *apischema.Record
		err := readJSON(r.Body, recordAPI)
		defer r.Body.Close()
		if err != nil {
			logger.Warnf("Failed to read JSON: %s", err.Error())
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}
		record := schemabuilder.BuildRecordFrom(recordAPI)
		err = ctx.ctrl.CreateRecord(record)
		if err != nil {
			logger.Warnf("Failed to get records from controller: %s", err.Error())
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}
		// TODO: get record from db
		writeJSONResponse(w, logger, schemabuilder.BuildRecordAPIFrom(record), http.StatusAccepted)
	}
}

func NewDeleteRecordHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "UpdateRecords"})
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		idStr := ps.ByName(IDParamName)
		idInt, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			logger.Warnf("Failed to convert path parameter id: %s", err.Error())
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}
		err = ctx.ctrl.DeleteRecord(idInt)
		if err != nil {
			logger.Warnf("Failed to get records from controller: %s", err.Error())
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func readJSON(requestBody io.ReadCloser, out any) error {
	// TODO: prevent overflow (read by batches or set max size)
	recordsJSON, err := io.ReadAll(requestBody)
	defer requestBody.Close()
	if err != nil {
		return err
	}
	err = json.Unmarshal(recordsJSON, &out)
	if err != nil {
		return err
	}
	return err
}

func writeJSONResponse(w http.ResponseWriter, logger log.Logger, body any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		logger.Warnf("Failed to write JSON response: %s", err.Error())
	}
	// TODO: do not log private info
	logger.Debugf("Response written: %+v", body)
}
