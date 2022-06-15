package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"

	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/controllerSchema"
)

const (
	IDParamName = "id"
)

type Controller interface {
	AllRecords() ([]controllerSchema.Record, error)
	Record(id uuid.UUID) (*controllerSchema.Record, error)
	CreateRecord(record *controllerSchema.Record) (*controllerSchema.Record, error)
	UpdateRecord(id uuid.UUID, record *controllerSchema.Record) (*controllerSchema.Record, error)
	DeleteRecord(id uuid.UUID) (*controllerSchema.Record, error)

	AllUsers() ([]controllerSchema.User, error)
	User(id uuid.UUID) (*controllerSchema.User, error)
	CreateUser(user *controllerSchema.User) (*controllerSchema.User, error)
	UpdateUser(id uuid.UUID, user *controllerSchema.User) (*controllerSchema.User, error)
	DeleteUser(id uuid.UUID) (*controllerSchema.User, error)
}

const (
	// Header Keys
	// Parse or get them from somewhere?
	CorrelationIDName = "X-Request-ID"
)

type RequestContext struct {
	corID uuid.UUID
	ps    httprouter.Params
}

type API struct {
	config *Config
	ctx    *APIContext
	server http.Server
}

type APIContext struct {
	controller Controller
	logger     log.Logger
}

type HandlerFunc func(rw http.ResponseWriter, r *http.Request, ctx *RequestContext)

func New(config *Config, controller Controller, logger log.Logger) *API {
	return &API{
		config: config,
		ctx: &APIContext{
			controller: controller,
			logger:     logger.WithFields(log.Fields{"service": "API"}),
		},
	}
}

func (api *API) Start() error {
	api.ctx.logger.Info("API started")
	router := httprouter.New()

	router.GET("/records", loggerMiddleware(api.ctx, AllRecordsHandler(api.ctx)))
	router.GET(fmt.Sprintf("/records/:%s", IDParamName), loggerMiddleware(api.ctx, RecordByIDHandler(api.ctx)))
	router.POST("/records", loggerMiddleware(api.ctx, CreateRecordHandler(api.ctx)))
	router.PUT(fmt.Sprintf("/records/:%s", IDParamName), loggerMiddleware(api.ctx, UpdateRecordHandler(api.ctx)))
	router.DELETE(fmt.Sprintf("/records/:%s", IDParamName), loggerMiddleware(api.ctx, DeleteRecordHandler(api.ctx)))

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

func loggerMiddleware(ctx *APIContext, next HandlerFunc) httprouter.Handle {
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
