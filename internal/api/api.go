package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"

	"github.com/okutsen/PasswordManager/schema/controllerSchema"

	"github.com/okutsen/PasswordManager/internal/log"
)

const (
	// PPN: Path Parameter Name
	IDPPN = "id"
	// HPN: Header Parameter Name
	CorrelationIDHPN      = "X-Request-ID"
	AuthorizationTokenHPN = "Authorization"

	RequestContextName    = "rctx"
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
	ctrl   Controller
	logger log.Logger
}

type HandlerFunc func(rw http.ResponseWriter, r *http.Request, ctx *RequestContext)

func New(config *Config, controller Controller, logger log.Logger) *API {
	return &API{
		config: config,
		ctx: &APIContext{
			ctrl: controller,
			logger:     logger.WithFields(log.Fields{"service": "API"}),
		},
	}
}

func (api *API) Start() error {
	api.ctx.logger.Info("API started")
	router := httprouter.New()

	router.GET("/records", ContextSetter(api.ctx.logger, NewListRecordsHandler(api.ctx)))
	router.POST("/records", ContextSetter(api.ctx.logger, NewCreateRecordHandler(api.ctx)))
	router.GET(fmt.Sprintf("/records/:%s", IDPPN), ContextSetter(api.ctx.logger, NewGetRecordHandler(api.ctx)))
	router.PUT(fmt.Sprintf("/records/:%s", IDPPN), ContextSetter(api.ctx.logger, NewUpdateRecordHandler(api.ctx)))
	router.DELETE(fmt.Sprintf("/records/:%s", IDPPN), ContextSetter(api.ctx.logger, NewDeleteRecordHandler(api.ctx)))

	router.GET("/records", ContextSetter(api.ctx.logger, NewListUsersHandler(api.ctx)))
	router.POST("/records", ContextSetter(api.ctx.logger, NewCreateUserHandler(api.ctx)))
	router.GET(fmt.Sprintf("/records/:%s", IDPPN), ContextSetter(api.ctx.logger, NewGetUserHandler(api.ctx)))
	router.PUT(fmt.Sprintf("/records/:%s", IDPPN), ContextSetter(api.ctx.logger, NewUpdateUserHandler(api.ctx)))
	router.DELETE(fmt.Sprintf("/records/:%s", IDPPN), ContextSetter(api.ctx.logger, NewDeleteUserHandler(api.ctx)))

	api.server = http.Server{Addr: api.config.Address(), Handler: router}

	return api.server.ListenAndServe()
}

func (api *API) Stop(ctx context.Context) error {
	api.ctx.logger.Infof("shutting down server")
	return api.server.Shutdown(ctx)
}
