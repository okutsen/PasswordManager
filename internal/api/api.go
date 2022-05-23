package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"

	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/apischema"
)

const (
	IDParamName = "id"
)

type RecordsController interface {
	AllRecords() ([]apischema.Record, error)
	Record(id uuid.UUID) (*apischema.Record, error)
	CreateRecord(record *apischema.Record) (*apischema.Record, error)
	UpdateRecord(id uuid.UUID, record *apischema.Record) (*apischema.Record, error)
	DeleteRecord(id uuid.UUID) error
}

type UsersController interface {
	AllUsers() ([]apischema.User, error)
	User(id uuid.UUID) (*apischema.User, error)
	CreateUser(user *apischema.User) (*apischema.User, error)
	UpdateUser(id uuid.UUID, user *apischema.User) (*apischema.User, error)
	DeleteUser(id uuid.UUID) error
}

type API struct {
	config *Config
	ctx    *APIContext
	server http.Server
}

type APIContext struct {
	recordsController RecordsController
	usersController   UsersController
	logger            log.Logger
}

func New(config *Config, ctrlRecords RecordsController, ctrlUsers UsersController, logger log.Logger) *API {
	return &API{
		config: config,
		ctx: &APIContext{
			recordsController: ctrlRecords,
			usersController:   ctrlUsers,
			logger:            logger.WithFields(log.Fields{"service": "API"}),
		},
	}
}

func (api *API) Start() error {
	api.ctx.logger.Info("API started")
	router := httprouter.New()

	router.GET("/records", loggerMiddleware(api.ctx, NewAllRecordsHandler(api.ctx)))
	router.GET(fmt.Sprintf("/records/:%s", IDParamName), loggerMiddleware(api.ctx, NewRecordByIDHandler(api.ctx)))
	router.POST("/records", loggerMiddleware(api.ctx, NewCreateRecordHandler(api.ctx)))
	router.PUT(fmt.Sprintf("/records/:%s", IDParamName), loggerMiddleware(api.ctx, NewUpdateRecordHandler(api.ctx)))
	router.DELETE(fmt.Sprintf("/records/:%s", IDParamName), loggerMiddleware(api.ctx, NewDeleteRecordHandler(api.ctx)))

	router.GET("/users", loggerMiddleware(api.ctx, AllUsersHandler(api.ctx)))
	router.GET(fmt.Sprintf("/users/:%s", IDParamName), loggerMiddleware(api.ctx, UserByIdHandler(api.ctx)))
	router.POST("/users", loggerMiddleware(api.ctx, CreateUserHandler(api.ctx)))
	router.PUT(fmt.Sprintf("/users/:%s", IDParamName), loggerMiddleware(api.ctx, UpdateUserHandler(api.ctx)))
	router.DELETE(fmt.Sprintf("/users/:%s", IDParamName), loggerMiddleware(api.ctx, DeleteUserHandler(api.ctx)))

	api.server = http.Server{Addr: api.config.Address(), Handler: router}

	return api.server.ListenAndServe()
}

func (api *API) Stop(ctx context.Context) error {
	api.ctx.logger.Infof("shutting down server")
	return api.server.Shutdown(ctx)
}

func loggerMiddleware(ctx *APIContext, handler httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx.logger.Infof("API: Endpoint Hit: %s %s%s", r.Method, r.Host, r.URL.Path)
		handler(rw, r, ps)
	}
}

func readJSON(requestBody io.ReadCloser, out any) error {
	// TODO: prevent overflow (read by batches or set max size)
	recordsJSON, err := io.ReadAll(requestBody)
	if err != nil {
		return err
	}
	defer requestBody.Close()

	err = json.Unmarshal(recordsJSON, &out)
	if err != nil {
		return err
	}

	fmt.Println(out)

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
