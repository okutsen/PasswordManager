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

func AllUsersHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "GetAllUsers"})
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		records, err := ctx.usersController.AllUsers()
		if err != nil {
			logger.Warnf("failed to get users: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, records, http.StatusOK)
	}
}

func UserByIdHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "GetUserByID"})
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		idStr := ps.ByName(IDParamName)
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			logger.Warnf("failed to convert path parameter id %s: %v", idStr, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidIDParam}, http.StatusBadRequest)
			return
		}

		user, err := ctx.usersController.User(id)
		if err != nil {
			logger.Warnf("failed to get user %d: %v", id, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, user, http.StatusOK)
	}
}

func CreateUserHandler(ctx *APIContext) httprouter.Handle {
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

		var user apischema.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			logger.Warnf("failed to unmarshal: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}

		// TODO: if exists return err (409 Conflict)
		createdUser, err := ctx.usersController.CreateUser(&user)
		if err != nil {
			logger.Warnf("failed to create user: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, createdUser, http.StatusAccepted)
	}
}

func UpdateUserHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "UpdateRecords"})
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		idStr := ps.ByName(IDParamName)
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			logger.Warnf("failed to convert path parameter id %s: %v", idStr, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidIDParam}, http.StatusBadRequest)
			return
		}

		var user apischema.User
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Warnf("failed to read body: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err = json.Unmarshal(body, &user)
		if err != nil {
			logger.Warnf("failed to unmarshal: %v", err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}

		updatedUser, err := ctx.usersController.UpdateUser(id, &user)
		if err != nil {
			logger.Warnf("failed to update user %d: %v", user.ID, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, logger, updatedUser, http.StatusAccepted)
	}
}

func DeleteUserHandler(ctx *APIContext) httprouter.Handle {
	logger := ctx.logger.WithFields(log.Fields{"handler": "DeleteUser"})
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		idStr := ps.ByName(IDParamName)
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			logger.Warnf("failed to convert path parameter id %s: %v", idStr, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InvalidJSONMessage}, http.StatusBadRequest)
			return
		}

		err = ctx.usersController.DeleteUser(id)
		if err != nil {
			logger.Warnf("failed to delete user %d: %v", id, err)
			writeJSONResponse(w, logger,
				apischema.Error{Message: apischema.InternalErrorMessage}, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
