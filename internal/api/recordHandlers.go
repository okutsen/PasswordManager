package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"

	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/apischema"
	"github.com/okutsen/PasswordManager/schema/schemabuilder"
)

func AllRecordsHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{"handler": "GetAllRecords"})
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger = logger.WithFields(log.Fields{
			"corID": rctx.corID,
		})
		controllerRecords, err := apictx.controller.AllRecords()
		if err != nil {
			logger.Warnf("failed to get records: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}
		APIRecord := schemabuilder.BuildAPIRecordsFromControllerRecords(controllerRecords)

		writeJSONResponse(w, logger, APIRecord, http.StatusOK)
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

		controllerRecord, err := apictx.controller.Record(id)
		if err != nil {
			logger.Warnf("failed to get record %s: %v", id, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}
		APIRecord := schemabuilder.BuildAPIRecordFromControllerRecord(controllerRecord)

		writeJSONResponse(w, logger, APIRecord, http.StatusOK)
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

		var apiRecord apischema.Record
		err = json.Unmarshal(body, &apiRecord)
		if err != nil {
			logger.Warnf("failed to unmarshal: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}
		controllerRecord := schemabuilder.BuildControllerRecordFromAPIRecord(&apiRecord)
		// TODO: if exists return err (409 Conflict)
		createRecord, err := apictx.controller.CreateRecord(&controllerRecord)
		if err != nil {
			logger.Warnf("failed to create record: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		apiRecord = schemabuilder.BuildAPIRecordFromControllerRecord(createRecord)
		writeJSONResponse(w, logger, apiRecord, http.StatusAccepted)
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

		controllerRecord := schemabuilder.BuildControllerRecordFromAPIRecord(&recordAPI)
		updatedRecord, err := apictx.controller.UpdateRecord(id, &controllerRecord)
		if err != nil {
			logger.Warnf("failed to update record %d: %v", recordAPI.ID, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		record := schemabuilder.BuildAPIRecordFromControllerRecord(updatedRecord)

		writeJSONResponse(w, logger, record, http.StatusAccepted)
	}
}

func DeleteRecordHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "DeleteRecord",
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

		controllerRecord, err := apictx.controller.DeleteRecord(id)
		if err != nil {
			logger.Warnf("failed to delete user %d: %v", id, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}
		record := schemabuilder.BuildAPIRecordFromControllerRecord(controllerRecord)

		writeJSONResponse(w, logger, record, http.StatusOK)
	}
}
