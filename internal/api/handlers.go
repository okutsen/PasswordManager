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

type HandlerContext struct {
	corID  uuid.UUID
	ctrl   Controller
	logger log.Logger
}

func NewHandlerContextFrom(ctx *APIContext, corID uuid.UUID) *HandlerContext {
	return &HandlerContext{
		corID: corID,
		ctrl:   ctx.ctrl,
		logger: ctx.logger.WithFields(log.Fields{"corID": corID.String()}),
	}
}

func NewHandlerContext(ctx *APIContext) *HandlerContext {
	// Use uuid.UUID or just string
	newCorID := uuid.New()
	ctx.logger.Debugf("Assigning new correlation id: %s", newCorID.String())
	return &HandlerContext{
		corID: newCorID,
		ctrl:   ctx.ctrl,
		logger: ctx.logger.WithFields(log.Fields{"corID": newCorID.String()}),
	}
}

// Name?
func NewInitMiddleware(ctx *APIContext, next HandlerCreator) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// When to Info vs Debug
		ctx.logger.Infof("Endpoint Hit: %s %s%s", r.Method, r.Host, r.URL.Path)
		corIDStr := r.Header.Get(CorrelationIDName)
		if corIDStr == "" {
			next(NewHandlerContext(ctx))(rw, r, ps)
			return
		}
		corID, err := uuid.Parse(corIDStr)
		if err != nil {
			// Error or Warn?
			ctx.logger.Warnf("Got invalid correlation id: %s", corIDStr)
			next(NewHandlerContext(ctx))(rw, r, ps)
			return
		}
		next(NewHandlerContextFrom(ctx, corID))(rw, r, ps)
	}
}

type HandlerCreator func(hctx *HandlerContext) httprouter.Handle

func NewGetAllRecordsHandler(hctx *HandlerContext) httprouter.Handle {
	logger := hctx.logger.WithFields(log.Fields{"handler": "GetAllRecords"})
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		records, err := hctx.ctrl.GetAllRecords()
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

func NewGetRecordHandler(hctx *HandlerContext) httprouter.Handle {
	logger := hctx.logger.WithFields(log.Fields{"handler": "GetRecord"})
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		idStr := ps.ByName(IDParamName)
		recordUUID, err := uuid.Parse(idStr)
		if err != nil {
			logger.Warnf("Failed to convert path parameter id: %s", err.Error())
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}
		record, err := hctx.ctrl.GetRecord(recordUUID)
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

func NewCreateRecordHandler(hctx *HandlerContext) httprouter.Handle {
	logger := hctx.logger.WithFields(log.Fields{"handler": "CreateRecords"})
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
		err = hctx.ctrl.CreateRecord(record)
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

func NewUpdateRecordHandler(hctx *HandlerContext) httprouter.Handle {
	logger := hctx.logger.WithFields(log.Fields{"handler": "UpdateRecords"})
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
		err = hctx.ctrl.CreateRecord(record)
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

func NewDeleteRecordHandler(hctx *HandlerContext) httprouter.Handle {
	logger := hctx.logger.WithFields(log.Fields{"handler": "UpdateRecords"})
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		idStr := ps.ByName(IDParamName)
		recordUUID, err := uuid.Parse(idStr)
		if err != nil {
			logger.Warnf("Failed to convert path parameter id: %s", err.Error())
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}
		err = hctx.ctrl.DeleteRecord(recordUUID)
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
