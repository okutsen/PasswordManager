package internal

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type ClientAPI struct {
	// TODO: log log.Logger
}

func NewClientAPI() *ClientAPI {
	// TODO: setup connection with server
	return &ClientAPI{}
}

func (c *ClientAPI) Start() {
	log.Info("Server started")
	router := httprouter.New()
	router.GET("/records", c.getRecords)
	router.GET("/records/:name", c.getRecords)
	router.POST("/records", c.createRecords)

	log.Fatal(http.ListenAndServe(":10000", router))
}

// TODO: pass params to model -> validate them and run query with filters
func (c *ClientAPI) getRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Info("Endpoint Hit: getRecords")
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
	log.Infof("Response writen %s", message)
	_, err := fmt.Fprintf(w, message)
	if err != nil {
		log.Error("getRecords: Failed to write responce")
	}
	// TODO: write JSON response
	// w.Header().Set("Content-Type", "application/json")
	// err := json.NewEncoder(w).Encode(config.Record)
	// if err != nil {
	// 	log.Printf("Failed to encode ports: %v\n -> Failed to write responce", err)
	// }
}

func (c *ClientAPI) createRecords(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	log.Info("Endpoint Hit: createRecords")
	w.WriteHeader(http.StatusAccepted)
	_, err := fmt.Fprintf(w, "New Record created")
	if err != nil {
		log.Error("createRecords: Failed to write responce")
	}

	// TODO: parse JSON input
}

func (c *ClientAPI) deleteRecords(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	log.Info("Endpoint Hit: deleteRecords")

	_, err := fmt.Fprintf(w, "Record deleted")
	if err != nil {
		log.Error("deleteRecords: Failed to write responce")
	}

	// TODO: parse JSON input
}
