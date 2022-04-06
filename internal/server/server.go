package server

import (
	"context"
	"net/http"

	"github.com/okutsen/PasswordManager/internal/log"
)

// FIXME: where to define?
var server_config = NewConfig()

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
			// TODO: add host to Addr: ServerHostURL + ":" + server_config.ServerListenPort
			Addr:         ":" + server_config.ServerListenPort,
			Handler:      NewHandler().router,
			ReadTimeout:  server_config.ReadTimeout,
			WriteTimeout: server_config.WriteTimeout,
		},
		handler: NewHandler(),
		log:     log.NewLogger(),
	}
}

func (d *Server) Start() {
	d.log.Print("Server started")

	if err := d.httpServer.ListenAndServe(); err != nil {
		_ = d.httpServer.Shutdown(context.TODO())
		d.log.Fatalf("Error: while running Server: %s", err.Error())
	}
}
