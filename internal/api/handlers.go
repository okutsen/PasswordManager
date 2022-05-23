package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"

	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/apischema"
	"github.com/okutsen/PasswordManager/schema/schemabuilder"
)

const (
	// Header Keys
	// Parse or get them from somewhere?
	CorrelationIDName = "X-Request-ID"
)

type RequestContext struct {
	corID uuid.UUID
	ps    httprouter.Params
}

type HandlerFunc func(rw http.ResponseWriter, r *http.Request, ctx *RequestContext)

// Name?
func InitMiddleware(ctx *APIContext, next HandlerFunc) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// When to Info vs Debug
		ctx.logger.Debugf("Endpoint Hit: %s %s%s", r.Method, r.Host, r.URL.Path)
		corIDStr := r.Header.Get(CorrelationIDName)
		corID := parseRequestID(corIDStr, ctx.logger)
		next(rw, r, &RequestContext{corID: corID, ps: ps})
	}
}

func parseRequestID(idStr string, logger log.Logger) uuid.UUID {
	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Warnf("Invalid corID <%s>: %s", idStr, err)
		newID := uuid.New()
		logger.Debugf("Setting new corID: %s", newID.String())
		return newID
	}
	return id
}

func NewGetAllRecordsHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{"handler": "GetAllRecords"})
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger = logger.WithFields(log.Fields{
			"corID": rctx.corID,
		})
		records, err := apictx.ctrl.GetAllRecords()
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

func NewGetRecordHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "GetRecord",
	})
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger := logger.WithFields(log.Fields{
			"corID": rctx.corID,
		})
		idStr := rctx.ps.ByName(IDParamName)
		recordUUID, err := uuid.Parse(idStr)
		if err != nil {
			logger.Warnf("Failed to convert path parameter id: %s", err.Error())
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}
		record, err := apictx.ctrl.GetRecord(recordUUID)
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

func NewCreateRecordHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "CreateRecord",
	})
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger := logger.WithFields(log.Fields{
			"corID": rctx.corID,
		})
		// TODO: check content type
		var recordAPI *apischema.Record
		err := readJSON(r.Body, &recordAPI)
		defer r.Body.Close()
		if err != nil {
			logger.Warnf("Failed to read JSON: %s", err.Error())
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}
		record := schemabuilder.BuildRecordFrom(recordAPI)
		// TODO: if exists return err (409 Conflict)
		err = apictx.ctrl.CreateRecord(record)
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

func NewUpdateRecordHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "UpdateRecords",
	})
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger := logger.WithFields(log.Fields{
			"corID": rctx.corID,
		})
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
		err = apictx.ctrl.CreateRecord(record)
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

func NewDeleteRecordHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "DeleteRecords",
	})
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger := logger.WithFields(log.Fields{
			"corID": rctx.corID,
		})
		idStr := rctx.ps.ByName(IDParamName)
		recordUUID, err := uuid.Parse(idStr)
		if err != nil {
			logger.Warnf("Failed to convert path parameter id: %s", err.Error())
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}
		err = apictx.ctrl.DeleteRecord(recordUUID)
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
	err = json.Unmarshal(recordsJSON, out)
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
