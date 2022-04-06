package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/okutsen/PasswordManager/internal/log"
)

// FIXME: where to define?
var api_config = NewConfig()

// log = log.NewLogger()

type API struct {
	serverConnection *http.Client
	log              log.Logger
}

func NewAPI() *API {
	return &API{
		serverConnection: &http.Client{
			Timeout: api_config.APIRequestTimeout,
		},
		log: log.NewLogger(),
	}
}

func (c *API) Start() {
	c.log.Print("API started")
	router := httprouter.New()
	// TODO: add route to config?
	router.GET("/records", c.getRecords)
	router.GET(fmt.Sprintf("/records/:%s", api_config.GetByIdParamName), c.getRecords)
	router.POST("/records", c.createRecords)

	// TODO: add host to Addr: APIHostURL + ":" + server_config.ServerListenPort
	c.log.Fatal(http.ListenAndServe(":"+api_config.APIListenPort, router))
}

// TODO: pass params to model -> validate them and run query with filters
func (c *API) getRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c.log.Print("API: Endpoint Hit: getRecords")
	// TODO: c.serverConnection.Get(config.DomainServerURL + ":" + config.DomainServerPort)
	idStr := ps.ByName(api_config.GetByIdParamName)
	if idStr == "" {
		responceBody := "Records:\n0,1,2,3,4,5"
		c.writeResponce(responceBody, http.StatusOK, w)
		return
	}
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		c.log.Print("Error: getRecords: %s", err.Error())
		c.writeResponce(api_config.ItemNotFoundMessage, http.StatusBadRequest, w)
		return
	}
	if 0 > idInt || idInt > 5 {
		c.writeResponce(api_config.ItemNotFoundMessage, http.StatusBadRequest, w)
		return
	}

	responceBody := fmt.Sprintf("Records:\n%d", idInt)
	c.writeResponce(responceBody, http.StatusOK, w)

	// TODO: write JSON response
}

func (c *API) writeResponce(body string, status int, w http.ResponseWriter) error {
	w.WriteHeader(status)
	_, err := fmt.Fprintf(w, body)
	if err != nil {
		c.log.Print("Error: getRecords: Failed to write responce")
		return err
	}
	c.log.Printf("Response writen\n%s", body)
	return nil
}

func (c *API) createRecords(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	c.log.Print("API: Endpoint Hit: createRecords")
	// TODO: c.serverConnection.Post(config.DomainServerURL + ":" + config.DomainServerPort, body)
	w.WriteHeader(http.StatusAccepted)
	_, err := fmt.Fprintf(w, "New Record created")
	if err != nil {
		c.log.Print("Error: createRecords: Failed to write responce")
	}

	// TODO: parse JSON input
}

func (c *API) deleteRecords(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	c.log.Print("Endpoint Hit: deleteRecords")
	// TODO: req := http.NewRequest(http.MethodDelete, config.DomainServerURL + ":" + config.DomainServerPort, body)
	// c.serverConnection.Do(req)
	_, err := fmt.Fprintf(w, "Record deleted")
	if err != nil {
		c.log.Print("Error: deleteRecords: Failed to write responce")
	}

	// TODO: parse JSON input
}
