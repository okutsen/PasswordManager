package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/okutsen/PasswordManager/internal/log"
)

// FIXME: where to define?
var apiConfig = NewConfig()

// log = log.NewLogger()

type API struct {
	serverConnection *http.Client
	log              log.Logger
}

func NewAPI() *API {
	return &API{
		serverConnection: &http.Client{
			Timeout: apiConfig.APIRequestTimeout,
		},
		log: log.NewLogger(),
	}
}

func (c *API) Start() {
	c.log.Print("API started")
	router := httprouter.New()
	// TODO: add route to config?
	router.GET("/records", c.getRecordsLogger(c.getRecords))
	router.GET(fmt.Sprintf("/records/:%s", apiConfig.GetByIdParamName), c.getRecordsLogger(c.getRecords))
	router.POST("/records", c.createRecordsLogger(c.createRecords))

	// TODO: add host to Addr: APIHostURL + ":" + server_config.ServerListenPort
	c.log.Fatal(http.ListenAndServe(":"+apiConfig.APIListenPort, router))
}

func (c *API) getRecordsLogger(handler httprouter.Handle) httprouter.Handle {
	loggedHandler := func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		c.log.Print("API: Endpoint Hit: getRecords")
		handler(rw, r, ps)
	}
	return loggedHandler
}

func (c *API) createRecordsLogger(handler httprouter.Handle) httprouter.Handle {
	loggedHandler := func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		c.log.Print("API: Endpoint Hit: createRecords")
		handler(rw, r, ps)
	}
	return loggedHandler
}

// TODO: pass params to model -> validate them and run query with filters
func (c *API) getRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// c.log.Print("API: Endpoint Hit: getRecords")
	// TODO: c.serverConnection.Get(config.DomainServerURL + ":" + config.DomainServerPort)
	idStr := ps.ByName(apiConfig.GetByIdParamName)
	if idStr == "" {
		responseBody := "Records:\n0,1,2,3,4,5"
		c.writeResponse(responseBody, http.StatusOK, w)
		return
	}
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		c.log.Print("Error: getRecords: %s", err.Error())
		c.writeResponse(apiConfig.ItemNotFoundMessage, http.StatusBadRequest, w)
		return
	}
	if 0 > idInt || idInt > 5 {
		c.writeResponse(apiConfig.ItemNotFoundMessage, http.StatusBadRequest, w)
		return
	}

	responseBody := fmt.Sprintf("Records:\n%d", idInt)
	c.writeResponse(responseBody, http.StatusOK, w)

	// TODO: write JSON response
}

func (c *API) writeResponse(body string, status int, w http.ResponseWriter) error {
	w.WriteHeader(status)
	_, err := fmt.Fprint(w, body)
	if err != nil {
		c.log.Print("Error: getRecords: Failed to write responce")
		return err
	}
	c.log.Printf("Response written\n%s", body)
	return nil
}

func (c *API) createRecords(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	c.log.Print("API: Endpoint Hit: createRecords")
	// TODO: c.serverConnection.Post(config.DomainServerURL + ":" + config.DomainServerPort, body)
	w.WriteHeader(http.StatusAccepted)
	_, err := fmt.Fprintf(w, "New Record created")
	if err != nil {
		c.log.Print("Error: createRecords: Failed to write response")
	}

	// TODO: parse JSON input
}

func (c *API) deleteRecords(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	c.log.Print("Endpoint Hit: deleteRecords")
	// TODO: req := http.NewRequest(http.MethodDelete, config.DomainServerURL + ":" + config.DomainServerPort, body)
	// c.serverConnection.Do(req)
	_, err := fmt.Fprintf(w, "Record deleted")
	if err != nil {
		c.log.Print("Error: deleteRecords: Failed to write response")
	}

	// TODO: parse JSON input
}
