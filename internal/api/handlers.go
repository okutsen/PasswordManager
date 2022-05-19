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
	corID  uuid.UUID
	ps     httprouter.Params
	*APIContext
}

func NewRequestContext(ctx *APIContext, ps httprouter.Params) *RequestContext {
	newCorID := uuid.New()
	ctx.logger.Debugf("Assigning new correlation id: %s", newCorID.String())
	return &RequestContext{
		corID:  newCorID,
		ps: ps,
		APIContext: ctx,
	}
}

func NewRequestContextFrom(ctx *APIContext, ps httprouter.Params, corID uuid.UUID) *RequestContext {
	ctx.logger.Debugf("Assigning correlation id: %s", corID.String())
	return &RequestContext{
		corID:  corID,
		ps: ps,
		APIContext: ctx,
	}
}

type InnerHandlerFunc func(rw http.ResponseWriter, r *http.Request, ctx *RequestContext)

// Name?
func InitMiddleware(ctx *APIContext, next InnerHandlerFunc) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// When to Info vs Debug
		ctx.logger.Infof("Endpoint Hit: %s %s%s", r.Method, r.Host, r.URL.Path)
		corIDStr := r.Header.Get(CorrelationIDName)
		if corIDStr == "" {
			next(rw, r, NewRequestContext(ctx, ps))
			return
		}
		corID, err := uuid.Parse(corIDStr)
		if err != nil {
			// Error or Warn?
			ctx.logger.Warnf("Got invalid correlation id: %s", corIDStr)
			next(rw, r, NewRequestContext(ctx, ps))
			return
		}
		next(rw, r, NewRequestContextFrom(ctx, ps, corID))
	}
}

func NewGetAllRecordsHandler() InnerHandlerFunc {
	// What should be done here again?
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger := rctx.logger.WithFields(log.Fields{
			"handler": "GetAllRecords", 
			"corID": rctx.corID,
		})
		records, err := rctx.ctrl.GetAllRecords()
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

func NewGetRecordHandler() InnerHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger := rctx.logger.WithFields(log.Fields{
			"handler": "GetRecord", 
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
		record, err := rctx.ctrl.GetRecord(recordUUID)
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

func NewCreateRecordHandler() InnerHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger := rctx.logger.WithFields(log.Fields{
			"handler": "CreateRecord", 
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
		err = rctx.ctrl.CreateRecord(record)
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

func NewUpdateRecordHandler() InnerHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger := rctx.logger.WithFields(log.Fields{
			"handler": "UpdateRecords", 
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
		err = rctx.ctrl.CreateRecord(record)
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

func NewDeleteRecordHandler() InnerHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger := rctx.logger.WithFields(log.Fields{
			"handler": "DeleteRecords", 
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
		err = rctx.ctrl.DeleteRecord(recordUUID)
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
