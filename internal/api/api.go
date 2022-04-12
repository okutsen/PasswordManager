package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/okutsen/PasswordManager/internal/log"
)

const (
	IDParamName    = "recordName"
	RecordNotFoundMessage = "Item not found"
	RecordCreatedMessage = "Record created"
)

type API struct {
	config *Config
	log  log.Logger
}

func New(config *Config, log log.Logger) *API {
	return &API{
		config: config,
		log:  log,
	}
}

func (c *API) endpointLogger(handler httprouter.Handle) httprouter.Handle {
	loggedHandler := func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		c.log.Infof("API: Endpoint Hit: %s %s%s\n", r.Host, r.URL.Path, r.Method)
		handler(rw, r, ps)
	}
	return loggedHandler
}

func (c *API) Start() error {
	c.log.Info("API started")
	router := httprouter.New()

	router.GET("/records", c.endpointLogger(c.getAllRecords))
	router.GET(fmt.Sprintf("/records/:%s", IDParamName), c.endpointLogger(c.getRecord))
	router.POST("/records", c.endpointLogger(c.createRecords))

	return http.ListenAndServe(c.config.Host+":"+c.config.Port, router)
}

func (c *API) getAllRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	responseBody := "Records:\n0,1,2,3,4,5"
	err := c.writeResponse(responseBody, http.StatusOK, w)
	c.log.Errorf("getAllRecords: Failed to write responce: %s", err.Error())

	// TODO: write JSON response
}

func (c *API) getRecord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName(IDParamName)
	// TODO: convert to correct type (uint)
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		c.log.Errorf("getRecord: Failed to convert path parameter id: %s", err.Error())
		c.writeResponse(RecordNotFoundMessage, http.StatusBadRequest, w)
		return
	}
	responseBody := fmt.Sprintf("Records:\n%d", idInt)
	err = c.writeResponse(responseBody, http.StatusOK, w)
	c.log.Errorf("getRecord: Failed to write responce: %s", err.Error())

	// TODO: write JSON response
}

func (c *API) createRecords(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := c.writeResponse(RecordCreatedMessage, http.StatusAccepted, w)
	c.log.Errorf("createRecords: Failed to write responce: %s", err.Error())

	// TODO: parse JSON input
}

func (c *API) writeResponse(body string, status int, w http.ResponseWriter) error {
	w.WriteHeader(status)
	_, err := fmt.Fprint(w, body)
	if err != nil {
		return err
	}
	c.log.Infof("Response written\n%s", body)
	return nil
}
