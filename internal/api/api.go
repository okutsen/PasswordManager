package api

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"

	"github.com/okutsen/PasswordManager/internal/log"
)

const (
	IDParamName           = "recordName"
	RecordNotFoundMessage = "Item not found"
	RecordCreatedMessage  = "Record created"
)

type API struct {
	config *Config
	log    log.Logger
	server http.Server
}

func New(config *Config, log log.Logger) *API {
	return &API{
		config: config,
		log:    log,
	}
}

func (api *API) endpointLogger(handler httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		api.log.Infof("API: Endpoint Hit: %s %s %s\n", r.Host, r.URL.Path, r.Method)
		handler(rw, r, ps)
	}
}

func (api *API) Start(_ context.Context) error {
	api.log.Info("server is starting")
	router := httprouter.New()
	api.log = api.log.WithFields(log.Fields{"service": "API"})

	router.GET("/records", api.endpointLogger(NewGetAllRecordsHandler(api.log)))
	router.GET(fmt.Sprintf("/records/:%s", IDParamName), api.endpointLogger(NewGetRecordHandler(api.log)))
	router.POST("/records", api.endpointLogger(NewCreateRecordsHandler(api.log)))

	api.server = http.Server{Addr: api.config.Address(), Handler: router}

	return api.server.ListenAndServe()
}

func (api *API) Stop(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	api.log.Info("shutting down server")
	return api.server.Shutdown(ctx)
}
