package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/okutsen/PasswordManager/internal/log"
)

const (
	IDParamName = "recordName"
)

type Controller interface {
	GetAllRecords() (string, error)
	GetRecord(uint64) (string, error)
	CreateRecords() (string, error)
}

type API struct {
	config *Config
	hctx   *HandlerContext
}

func New(config *Config, ctrl Controller, logger log.Logger) *API {
	return &API{
		config: config,
		hctx: &HandlerContext{
			ctrl:  ctrl,
			logger: logger.WithFields(log.Fields{"service": "API"}),
		},
	}
}

func (api *API) endpointLogger(handler httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		api.hctx.logger.Infof("API: Endpoint Hit: %s %s%s\n", r.Host, r.URL.Path, r.Method)
		handler(rw, r, ps)
	}
}

func (api *API) Start() error {
	api.hctx.logger.Info("API started")
	router := httprouter.New()

	router.GET("/records", api.endpointLogger(NewGetAllRecordsHandler(api.hctx)))
	router.GET(fmt.Sprintf("/records/:%s", IDParamName), api.endpointLogger(NewGetRecordHandler(api.hctx)))
	router.POST("/records", api.endpointLogger(NewCreateRecordsHandler(api.hctx)))

	return http.ListenAndServe(api.config.Address(), router)
}
