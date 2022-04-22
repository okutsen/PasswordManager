package api

import (
	"fmt"
	"net/http"

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
}

func New(config *Config, log log.Logger) *API {
	return &API{
		config: config,
		log:    log,
	}
}

func (c *API) endpointLogger(handler httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		c.log.Infof("API: Endpoint Hit: %s %s%s\n", r.Host, r.URL.Path, r.Method)
		handler(rw, r, ps)
	}
}

func (api *API) Start() error {
	api.log.Info("API started")
	router := httprouter.New()
	api.log = api.log.WithFields(log.Fields{"service": "API"})

	router.GET("/records", api.endpointLogger(NewGetAllRecordsHandler(api.log)))
	router.GET(fmt.Sprintf("/records/:%s", IDParamName), api.endpointLogger(NewGetRecordHandler(api.log)))
	router.POST("/records", api.endpointLogger(NewCreateRecordsHandler(api.log)))

	return http.ListenAndServe(api.config.Address(), router)
}
