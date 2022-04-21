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
		// TODO: create context logger, get correlationID
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

	return http.ListenAndServe(api.config.Address(), router)
}

func (api *API) getAllRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	contextLogger := api.log.WithFields(log.Fields{"handler": "getAllRecords"})

	// TODO: get from controller
	responseBody := "Records:\n0,1,2,3,4,5"

	writeResponse(responseBody, http.StatusOK, w, contextLogger)

	// TODO: write JSON response
}

func (api *API) getRecord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	contextLogger := api.log.WithFields(log.Fields{"handler": "getRecord"})
	idStr := ps.ByName(IDParamName)
	// TODO: convert to correct type (uint)
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		contextLogger.Warnf("Failed to convert path parameter id: %s", err.Error())
		writeResponse(RecordNotFoundMessage, http.StatusBadRequest, w, contextLogger)
		return
	}
	responseBody := fmt.Sprintf("Records:\n%d", idInt)
	writeResponse(responseBody, http.StatusOK, w, contextLogger)

	// TODO: write JSON response
}

func (api *API) createRecords(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	contextLogger := api.log.WithFields(log.Fields{"handler": "createRecords"})
	writeResponse(RecordCreatedMessage, http.StatusAccepted, w, contextLogger)
	// TODO: parse JSON input
}

func writeResponse(body string, status int, w http.ResponseWriter, logger log.Logger) {
	w.WriteHeader(status)
	_, err := fmt.Fprint(w, body)
	if err != nil {
		logger.Warnf("Failed to write response: %s", err.Error())
	}
	logger.Infof("Response written\n%s", body)
}
