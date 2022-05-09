package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/okutsen/PasswordManager/domain"
	"github.com/okutsen/PasswordManager/internal/log"
)

const (
	IDParamName = "recordName"
)

type Controller interface {
	GetAllRecords() ([]domain.Record, error)
	GetRecord(uint64) ([]domain.Record, error)
	CreateRecords([]domain.Record) error
}

type API struct {
	config *Config
	hctx   *HandlerContext
	server http.Server
}

func New(config *Config, ctrl Controller, logger log.Logger) *API {
	return &API{
		config: config,
		hctx: &HandlerContext{
			ctrl:   ctrl,
			logger: logger.WithFields(log.Fields{"service": "API"}),
		},
	}
}

func (api *API) endpointLogger(handler httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		api.hctx.logger.Infof("API: Endpoint Hit: %s %s%s", r.Method, r.Host, r.URL.Path)
		handler(rw, r, ps)
	}
}

func (api *API) Start() error {
	api.hctx.logger.Info("server is starting")
	router := httprouter.New()

	router.GET("/records", api.endpointLogger(NewGetAllRecordsHandler(api.hctx)))
	router.GET(fmt.Sprintf("/records/:%s", IDParamName), api.endpointLogger(NewGetRecordHandler(api.hctx)))
	router.POST("/records", api.endpointLogger(NewCreateRecordsHandler(api.hctx)))

	api.server = http.Server{Addr: api.config.Address(), Handler: router}

	return api.server.ListenAndServe()
}

func (api *API) Stop(ctx context.Context) error {
	api.hctx.logger.Infof("shutting down server")
	return api.server.Shutdown(ctx)
}
