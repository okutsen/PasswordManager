package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/apischema"
)

const (
	RecordCreatedMessage = "Record created"
)

func NewAllRecordsHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "GetAllRecords"})
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		records, err := ctx.recordsController.AllRecords()
		if err != nil {
			logger.Warnf("failed to get records: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, records, http.StatusOK)
	}
}

func NewRecordByIDHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "GetRecord"})
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		idStr := ps.ByName(IDParamName)
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			logger.Warnf("failed to convert path parameter id %s: %v", idStr, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}

		record, err := ctx.recordsController.Record(id)
		if err != nil {
			logger.Warnf("failed to get record %d: %v", id, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, record, http.StatusOK)
	}
}

func NewCreateRecordHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "CreateRecords"})
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
		record, err := ctx.recordsController.CreateRecord(&recordAPI)
		if err != nil {
			logger.Warnf("failed to create record: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, record, http.StatusAccepted)
	}
}

func NewUpdateRecordHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "UpdateRecords"})
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		idStr := ps.ByName(IDParamName)
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			logger.Warnf("failed to convert path parameter id %s: %v", idStr, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
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

		updatedRecord, err := ctx.recordsController.UpdateRecord(id, &recordAPI)
		if err != nil {
			logger.Warnf("failed to update record %d: %v", recordAPI.ID, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, updatedRecord, http.StatusAccepted)
	}
}

func NewDeleteRecordHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "DeleteRecord"})
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		idStr := ps.ByName(IDParamName)
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			logger.Warnf("failed to convert path parameter id %s: %v", idStr, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}

		err = ctx.recordsController.DeleteRecord(id)
		if err != nil {
			logger.Warnf("failed to delete record %d: %v", id, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
