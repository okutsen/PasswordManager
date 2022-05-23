package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"

	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/apischema"
)

func AllRecordsHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{"handler": "GetAllRecords"})
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger = logger.WithFields(log.Fields{
			"corID": rctx.corID,
		})
		records, err := apictx.controller.AllRecords()
		if err != nil {
			logger.Warnf("failed to get records: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, records, http.StatusOK)
	}
}

func RecordByIDHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "GetRecord",
	})
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger := logger.WithFields(log.Fields{
			"corID": rctx.corID,
		})
		idStr := rctx.ps.ByName(IDParamName)
		id, err := uuid.Parse(idStr)
		if err != nil {
			logger.Warnf("failed to convert path parameter id %s: %v", idStr, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidIDParam}, http.StatusBadRequest)
			return
		}

		record, err := apictx.controller.Record(id)
		if err != nil {
			logger.Warnf("failed to get record %s: %v", id, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, record, http.StatusOK)
	}
}

func CreateRecordHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "CreateRecord",
	})
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger := logger.WithFields(log.Fields{
			"corID": rctx.corID,
		})
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Warnf("failed to read body: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var recordAPI apischema.Record
		err = json.Unmarshal(body, &recordAPI)
		if err != nil {
			logger.Warnf("failed to unmarshal: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}

		// TODO: if exists return err (409 Conflict)
		record, err := apictx.controller.CreateRecord(&recordAPI)
		if err != nil {
			logger.Warnf("failed to create record: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, record, http.StatusAccepted)
	}
}

func UpdateRecordHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "UpdateRecords",
	})
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger := logger.WithFields(log.Fields{
			"corID": rctx.corID,
		})

		idStr := rctx.ps.ByName(IDParamName)
		id, err := uuid.Parse(idStr)
		if err != nil {
			logger.Warnf("failed to convert path parameter id %s: %v", idStr, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidIDParam}, http.StatusBadRequest)
			return
		}

		var recordAPI apischema.Record

		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Warnf("failed to read body: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err = json.Unmarshal(body, &recordAPI)
		if err != nil {
			logger.Warnf("failed to unmarshal: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}

		updatedRecord, err := apictx.controller.UpdateRecord(id, &recordAPI)
		if err != nil {
			logger.Warnf("failed to update record %d: %v", recordAPI.ID, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, updatedRecord, http.StatusAccepted)
	}
}

func DeleteRecordHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "DeleteRecords",
	})
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger := logger.WithFields(log.Fields{
			"corID": rctx.corID,
		})
		idStr := rctx.ps.ByName(IDParamName)
		id, err := uuid.Parse(idStr)
		if err != nil {
			logger.Warnf("failed to convert path parameter id %s: %v", idStr, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidIDParam}, http.StatusBadRequest)
			return
		}

		_, err = apictx.controller.DeleteRecord(id)
		if err != nil {
			logger.Warnf("failed to delete record %d: %v", id, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
