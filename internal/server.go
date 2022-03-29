package internal

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	// log Logger
}

func (c *Client) getRecords(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	log.Info("Endpoint Hit: getRecords")
	var message string = "Records:\n1,2,3,4"
	log.Infof("Response writen %s", message)
	fmt.Fprintf(w, message)

	// TODO: write JSON response
	// w.Header().Set("Content-Type", "application/json")
	// err := json.NewEncoder(w).Encode(config.Record)
	// if err != nil {
	// 	log.Printf("Failed to encode ports: %v\n -> Failed to write responce", err)
	// }
}

func (c *Client) createRecords(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	log.Info("Endpoint Hit: getRecords")

	fmt.Fprintf(w, "New Record created")

	// TODO: parse JSON input
}

func (c *Client) Start() {
	log.Info("Server started")
	router := httprouter.New()
	router.GET("/records", c.getRecords)
	router.POST("/records", c.createRecords)

	log.Fatal(http.ListenAndServe(":10000", router))
}
