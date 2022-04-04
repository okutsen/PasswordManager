package internal

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/okutsen/PasswordManager/config"
)

const getByIdParamName = "recordName"

type ClientAPI struct {
	domainConnection *http.Client
	log              Logger
}

func NewClientAPI() *ClientAPI {
	// TODO: setup connection with server
	return &ClientAPI{
		domainConnection: &http.Client{
			Timeout: 60 * time.Second,
		},
		log: NewLogger(),
	}
}

func (c *ClientAPI) Start() {
	c.log.Print("ClientAPI started")
	router := httprouter.New()
	router.GET("/records", c.getRecords)
	router.GET(fmt.Sprintf("/records/:%s", getByIdParamName), c.getRecords)
	router.POST("/records", c.createRecords)

	c.log.Fatal(http.ListenAndServe(":"+config.ClientAPIPort, router))
}

// TODO: pass params to model -> validate them and run query with filters
func (c *ClientAPI) getRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c.log.Print("ClientAPI: Endpoint Hit: getRecords")
	// TODO: c.domainConnection.Get(config.DomainServerURL + ":" + config.DomainServerPort)
	var message string
	recordName := ps.ByName("recordName")
	if recordName != "" {
		recordName, err := strconv.Atoi(recordName)
		if err != nil && 0 > recordName || recordName > 5 {
			message = "Record not found"
			w.WriteHeader(http.StatusBadRequest)
		} else {
			message = fmt.Sprintf("Records:\n%d", recordName)
			w.WriteHeader(http.StatusOK)
		}
	} else {
		message = "Records:\n0,1,2,3,4,5"
		w.WriteHeader(http.StatusOK)
	}
	c.log.Printf("Response writen %s", message)
	_, err := fmt.Fprintf(w, message)
	if err != nil {
		c.log.Print("Error: getRecords: Failed to write responce")
	}
	// TODO: write JSON response
	// w.Header().Set("Content-Type", "application/json")
	// err := json.NewEncoder(w).Encode(config.Record)
	// if err != nil {
	// 	log.Printf("Failed to encode ports: %v\n -> Failed to write responce", err)
	// }
}

func (c *ClientAPI) createRecords(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	c.log.Print("ClientAPI: Endpoint Hit: createRecords")
	// TODO: c.domainConnection.Post(config.DomainServerURL + ":" + config.DomainServerPort, body)
	w.WriteHeader(http.StatusAccepted)
	_, err := fmt.Fprintf(w, "New Record created")
	if err != nil {
		c.log.Print("Error: createRecords: Failed to write responce")
	}

	// TODO: parse JSON input
}

func (c *ClientAPI) deleteRecords(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	c.log.Print("Endpoint Hit: deleteRecords")
	// TODO: req := http.NewRequest(http.MethodDelete, config.DomainServerURL + ":" + config.DomainServerPort, body)
	// c.domainConnection.Do(req)
	_, err := fmt.Fprintf(w, "Record deleted")
	if err != nil {
		c.log.Print("Error: deleteRecords: Failed to write responce")
	}

	// TODO: parse JSON input
}
