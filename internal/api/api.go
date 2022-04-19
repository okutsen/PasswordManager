package api

import (
	"fmt"
	"net/http"
	"strconv"

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

func (api *API) endpointLogger(handler httprouter.Handle) httprouter.Handle {
	loggedHandler := func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		api.log.Infof("API: Endpoint Hit: %s %s%s\n", r.Host, r.URL.Path, r.Method)
		handler(rw, r, ps)
	}
	return loggedHandler
}

func (api *API) Start() error {
	api.log.Info("API started")
	router := httprouter.New()

	router.GET("/records", api.endpointLogger(api.getAllRecords))
	router.GET(fmt.Sprintf("/records/:%s", IDParamName), api.endpointLogger(api.getRecord))
	router.POST("/records", api.endpointLogger(api.createRecords))

	return http.ListenAndServe(api.config.Addr, router)
}

func (api *API) getAllRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// contextLogger := api.
	responseBody := "Records:\n0,1,2,3,4,5"
	err := api.writeResponse(responseBody, http.StatusOK, w)
	if err != nil {
		api.log.Errorf("getAllRecords: Failed to write response: %s", err.Error())
	}

	// TODO: write JSON response
}

func (api *API) getRecord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName(IDParamName)
	// TODO: convert to correct type (uint)
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		api.log.Errorf("getRecord: Failed to convert path parameter id: %s", err.Error())
		err := api.writeResponse(RecordNotFoundMessage, http.StatusBadRequest, w)
		if err != nil {
			api.log.Errorf("getRecord: Failed to write response: %s", err.Error())
		}
		return
	}
	responseBody := fmt.Sprintf("Records:\n%d", idInt)
	err = api.writeResponse(responseBody, http.StatusOK, w)
	if err != nil {
		api.log.Errorf("getRecord: Failed to write response: %s", err.Error())
	}

	// TODO: write JSON response
}

func (c *API) createRecords(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := c.writeResponse(RecordCreatedMessage, http.StatusAccepted, w)
	if err != nil {
		c.log.Errorf("createRecords: Failed to write response: %s", err.Error())
	}

	// TODO: parse JSON input
}

func (api *API) writeResponse(body string, status int, w http.ResponseWriter) error {
	w.WriteHeader(status)
	_, err := fmt.Fprint(w, body)
	if err != nil {
		return err
	}
	api.log.Infof("Response written\n%s", body)
	return nil
}
