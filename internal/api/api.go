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
	port string
	log  log.Logger
}

func New(config *Config, log log.Logger) *API {
	return &API{
		port: config.APIListenPort,
		log:  log,
	}
}

func (c *API) Start() error {
	c.log.Info("API started")
	router := httprouter.New()

	router.GET("/records", c.endpointLogger(c.getAllRecords))
	router.GET(fmt.Sprintf("/records/:%s", IDParamName), c.endpointLogger(c.getRecord))
	router.POST("/records", c.endpointLogger(c.createRecords))

	return http.ListenAndServe(":"+c.port, router)
}

func (c *API) endpointLogger(handler httprouter.Handle) httprouter.Handle {
	loggedHandler := func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		c.log.Infof("API: Endpoint Hit: %s%s with method %s\n", r.Host, r.URL.Path, r.Method)
		handler(rw, r, ps)
	}
	return loggedHandler
}

func (c *API) getAllRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	responseBody := "Records:\n0,1,2,3,4,5"
	c.writeResponse(responseBody, http.StatusOK, w)

	// TODO: write JSON response
}

func (c *API) getRecord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName(IDParamName)
	// TODO: convert to correct type (uint)
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		c.log.Errorf("getRecords: %s", err.Error())
		c.writeResponse(RecordNotFoundMessage, http.StatusBadRequest, w)
		return
	}
	responseBody := fmt.Sprintf("Records:\n%d", idInt)
	c.writeResponse(responseBody, http.StatusOK, w)

	// TODO: write JSON response
}

func (c *API) createRecords(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	c.writeResponse(RecordCreatedMessage, http.StatusAccepted, w)

	// TODO: parse JSON input
}

func (c *API) writeResponse(body string, status int, w http.ResponseWriter) error {
	w.WriteHeader(status)
	_, err := fmt.Fprint(w, body)
	if err != nil {
		c.log.Error("Failed to write responce")
		return err
	}
	c.log.Infof("Response written\n%s", body)
	return nil
}
