package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"

	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/apischema"
)

func AllUsersHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{"handler": "GetAllUsers"})
	return func(w http.ResponseWriter, r *http.Request, rctx *RequestContext) {
		logger = logger.WithFields(log.Fields{
			"corID": rctx.corID,
		})
		users, err := apictx.controller.AllUsers()
		if err != nil {
			logger.Warnf("failed to get users: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, users, http.StatusOK)
	}
}

func UserByIdHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "GetUser",
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

		user, err := apictx.controller.User(id)
		if err != nil {
			logger.Warnf("failed to get user %s: %v", id, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, user, http.StatusOK)
	}
}
func CreateUserHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "CreateUser",
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

		var userAPI apischema.User
		err = json.Unmarshal(body, &userAPI)
		if err != nil {
			logger.Warnf("failed to unmarshal: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}

		// TODO: if exists return err (409 Conflict)
		user, err := apictx.controller.CreateUser(&userAPI)
		if err != nil {
			logger.Warnf("failed to create user: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, user, http.StatusAccepted)
	}
}

func UpdateUserHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "UpdateUser",
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

		var userAPI apischema.User

		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Warnf("failed to read body: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err = json.Unmarshal(body, &userAPI)
		if err != nil {
			logger.Warnf("failed to unmarshal: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}

		updateUser, err := apictx.controller.UpdateUser(id, &userAPI)
		if err != nil {
			logger.Warnf("failed to update user %d: %v", userAPI.ID, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, updateUser, http.StatusAccepted)
	}
}

func DeleteUserHandler(apictx *APIContext) HandlerFunc {
	logger := apictx.logger.WithFields(log.Fields{
		"handler": "DeleteUser",
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

		user, err := apictx.controller.DeleteUser(id)
		if err != nil {
			logger.Warnf("failed to delete user %d: %v", id, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, user, http.StatusOK)
	}
}
