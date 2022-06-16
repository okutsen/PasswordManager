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
	// PPN: Path Parameter Name
	RecordIDPPN = "RecordID"
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

	router.GET("/records", ContextSetter(api.ctx.logger, NewListRecordsHandler(api.ctx)))
	router.POST("/records", ContextSetter(api.ctx.logger, NewCreateRecordHandler(api.ctx)))
	router.GET(fmt.Sprintf("/records/:%s", RecordIDPPN), ContextSetter(api.ctx.logger, NewGetRecordHandler(api.ctx)))
	router.PUT(fmt.Sprintf("/records/:%s", RecordIDPPN), ContextSetter(api.ctx.logger, NewUpdateRecordHandler(api.ctx)))
	router.DELETE(fmt.Sprintf("/records/:%s", RecordIDPPN), ContextSetter(api.ctx.logger, NewDeleteRecordHandler(api.ctx)))

	api.server = http.Server{Addr: api.config.Address(), Handler: router}

	return api.server.ListenAndServe()
}

func (api *API) Stop(ctx context.Context) error {
	api.ctx.logger.Infof("Shutting down server")
	return api.server.Shutdown(ctx)
}
