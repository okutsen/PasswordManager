package internal

import (
	"context"
	"net/http"
	"time"

	"github.com/okutsen/PasswordManager/config"
)

type DomainServer struct {
	httpServer *http.Server
	// TODO: add dependency injection (create Handler interface)
	handler *Handler
	log     Logger
}

func NewDomainServer() *DomainServer {

	return &DomainServer{
		httpServer: &http.Server{
			Addr:         ":" + config.DomainServerPort,
			Handler:      NewHandler().router,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		handler: NewHandler(),
		log:     NewLogger(),
	}
}

func (d *DomainServer) Start() {
	d.log.Print("Server started")

	if err := d.httpServer.ListenAndServe(); err != nil {
		_ = d.httpServer.Shutdown(context.TODO())
		d.log.Fatalf("Error: while running DomainServer: %s", err.Error())
	}
}
