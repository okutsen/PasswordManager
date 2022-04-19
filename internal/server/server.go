package server

import (
	"context"
	"net/http"

	"github.com/okutsen/PasswordManager/internal/log"
)

// FIXME: where to define?
var serverConfig = NewConfig()

// log = log.NewLogger()

type Server struct {
	httpServer *http.Server
	// TODO: add dependency injection (create Handler interface)
	handler *Handler
	log     log.Logger
}

func NewServer() *Server {

	return &Server{
		httpServer: &http.Server{
			// TODO: add host to Addr: ServerHostURL + ":" + serverConfig.ServerListenPort
			Addr:         ":" + serverConfig.ServerListenPort,
			Handler:      NewHandler().router,
			ReadTimeout:  serverConfig.ReadTimeout,
			WriteTimeout: serverConfig.WriteTimeout,
		},
		handler: NewHandler(),
	}
}

func (d *Server) Start() {
	d.log.Info("Server started")

	if err := d.httpServer.ListenAndServe(); err != nil {
		_ = d.httpServer.Shutdown(context.TODO())
		d.log.Fatalf("Error: while running Server: %s", err.Error())
	}
}
