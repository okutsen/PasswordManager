package api

import (
	"context"
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
	// HPN: Header Parameter Name
	CorrelationIDHPN      = "X-Request-ID"
	AuthorizationTokenHPN = "Authorization"
	RequestContextName    = "rctx"
)

type RequestContext struct {
	corID uuid.UUID
	ps    httprouter.Params
}

// unpackContext gets and validates RequestContext from ctx
func unpackRequestContext(ctx context.Context, logger log.Logger) *RequestContext {
	rctx, ok := ctx.Value(RequestContextName).(*RequestContext)
	if !ok {
		logger.Fatalf("Failed to unpack request context, got: %s", rctx)
	}
	return rctx
}

// getRecordID checks if id is set and returns the result of uuid parsing
func getRecordID(ps httprouter.Params, logger log.Logger) (uuid.UUID, error) {
	idStr := ps.ByName(RecordIDPPN)
	if idStr == "" {
		logger.Fatal("Failed to get path parameter: there is no record id")
	}
	return uuid.Parse(idStr)
}

func NewListRecordsHandler(apictx *APIContext) http.HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{"handler": "GetAllRecords"})
	return func(w http.ResponseWriter, r *http.Request) {
		rctx := unpackRequestContext(r.Context(), logger)
		logger = logger.WithFields(log.Fields{
			"cor_id": rctx.corID.String(),
		})
		records, err := apictx.ctrl.ListRecords()
		if err != nil {
			logger.Warnf("Failed to get records from controller: %s", err.Error())
			writeResponse(w,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError, logger)
			return
		}
		recordsAPI := schemabuilder.BuildRecordsAPIFrom(records)
		// Write JSON by stream?
		writeResponse(w, recordsAPI, http.StatusOK, logger)
	}
}

func NewGetRecordHandler(apictx *APIContext) http.HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "GetRecord",
	})
	return func(w http.ResponseWriter, r *http.Request) {
		rctx := unpackRequestContext(r.Context(), logger)
		logger = logger.WithFields(log.Fields{
			"cor_id": rctx.corID.String(),
		})
		recordID, err := getRecordID(rctx.ps, logger)
		if err != nil {
			logger.Warnf("Invalid record id: %s", err.Error())
			writeResponse(w, 
				apischema.Error{Message: apischema.InvalidRecordIDMessage}, http.StatusBadRequest, logger)
			return
		}
		record, err := apictx.ctrl.GetRecord(recordID)
		if err != nil {
			logger.Warnf("Failed to get records from controller: %s", err.Error())
			writeResponse(w,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError, logger)
			return
		}
		// TODO: get record from db
		writeResponse(w, schemabuilder.BuildRecordAPIFrom(record), http.StatusOK, logger)
	}
}

func NewCreateRecordHandler(apictx *APIContext) http.HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "CreateRecord",
	})
	return func(w http.ResponseWriter, r *http.Request) {
		rctx := unpackRequestContext(r.Context(), logger)
		logger = logger.WithFields(log.Fields{
			"cor_id": rctx.corID.String(),
		})
		var recordAPI *apischema.Record
		err := readJSON(r.Body, &recordAPI)
		defer r.Body.Close()
		if err != nil {
			logger.Warnf("Failed to read JSON: %s", err.Error())
			writeResponse(w,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest, logger)
			return
		}
		record := schemabuilder.BuildRecordFrom(recordAPI)
		// TODO: if exists return err (409 Conflict)
		// FIXME: return created struct
		err = apictx.ctrl.CreateRecord(record)
		if err != nil {
			logger.Warnf("Failed to get records from controller: %s", err.Error())
			writeResponse(w,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError, logger)
			return
		}
		// TODO: get record from db
		writeResponse(w, schemabuilder.BuildRecordAPIFrom(record), http.StatusCreated, logger)
	}
}

func NewUpdateRecordHandler(apictx *APIContext) http.HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "UpdateRecords",
	})
	return func(w http.ResponseWriter, r *http.Request) {
		rctx := unpackRequestContext(r.Context(), logger)
		logger = logger.WithFields(log.Fields{
			"cor_id": rctx.corID.String(),
		})
		recordID, err := getRecordID(rctx.ps, logger)
		if err != nil {
			logger.Warnf("Invalid record id: %s", err.Error())
			writeResponse(w, 
				apischema.Error{Message: apischema.InvalidRecordIDMessage}, http.StatusBadRequest, logger)
			return
		}
		var recordAPI *apischema.Record
		err = readJSON(r.Body, recordAPI)
		defer r.Body.Close()
		if err != nil {
			logger.Warnf("Failed to read JSON: %s", err.Error())
			writeResponse(w,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest, logger)
			return
		}
		if recordID != recordAPI.ID{
			logger.Warn("Record id from path parameter doesn't match id from new record structure")
			writeResponse(w,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest, logger)
			return
		}
		record := schemabuilder.BuildRecordFrom(recordAPI)
		err = apictx.ctrl.UpdateRecord(record)
		if err != nil {
			logger.Warnf("Failed to get records from controller: %s", err.Error())
			writeResponse(w,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError, logger)
			return
		}
		// TODO: get record from db
		writeResponse(w, 
			schemabuilder.BuildRecordAPIFrom(record), http.StatusAccepted, logger)
	}
}

func NewDeleteRecordHandler(apictx *APIContext) http.HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "DeleteRecords",
	})
	return func(w http.ResponseWriter, r *http.Request) {
		rctx := unpackRequestContext(r.Context(), logger)
		logger = logger.WithFields(log.Fields{
			"cor_id": rctx.corID.String(),
		})
		recordID, err := getRecordID(rctx.ps, logger)
		if err != nil {
			logger.Warnf("Invalid record id: %s", err.Error())
			writeResponse(w, 
				apischema.Error{Message: apischema.InvalidRecordIDMessage}, http.StatusBadRequest, logger)
			return
		}
		err = apictx.ctrl.DeleteRecord(recordID)
		if err != nil {
			logger.Errorf("Failed to get records from controller: %s", err.Error())
			writeResponse(w,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError, logger)
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

func writeResponse(w http.ResponseWriter, body any, statusCode int, logger log.Logger) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		logger.Warnf("Failed to write JSON response: %s", err.Error())
	}
	// TODO: do not log private info
	logger.Debugf("Response written: %+v", body)
}
