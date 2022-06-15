package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"

	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/dbschema"
)

const (
	IDPathParamName = "RecordID"
)

type Controller interface {
	ListRecords() ([]*dbschema.Record, error)
	GetRecord(uuid.UUID) (*dbschema.Record, error)
	CreateRecord(*dbschema.Record) error
	UpdateRecord(*dbschema.Record) error
	DeleteRecord(uuid.UUID) error
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

func New(config *Config, ctrl Controller, logger log.Logger) *API {
	return &API{
		config: config,
		ctx: &APIContext{
			ctrl:   ctrl,
			logger: logger.WithFields(log.Fields{"service": "API"}),
		},
	}
}

func (api *API) Start() error {
	api.ctx.logger.Info("API started")
	router := httprouter.New()

	router.GET("/records", ContextSetter(api.ctx, NewListRecordsHandler(api.ctx)))
	router.GET(fmt.Sprintf("/records/:%s", IDPathParamName), ContextSetter(api.ctx, NewGetRecordHandler(api.ctx)))
	router.POST("/records", ContextSetter(api.ctx, NewCreateRecordHandler(api.ctx)))
	router.PUT("/records", ContextSetter(api.ctx, NewUpdateRecordHandler(api.ctx)))
	router.DELETE(fmt.Sprintf("/records/:%s", IDPathParamName), ContextSetter(api.ctx, NewDeleteRecordHandler(api.ctx)))

	api.server = http.Server{Addr: api.config.Address(), Handler: router}

	return api.server.ListenAndServe()
}

func (api *API) Stop(ctx context.Context) error {
	api.ctx.logger.Infof("Shutting down server")
	return api.server.Shutdown(ctx)
}
