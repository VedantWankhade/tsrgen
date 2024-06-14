package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vedantwankhade/tsrgen/gateway-service/internal/ports"
)

type server struct {
	app ports.APIPort
}

func NewServer(app ports.APIPort) *server {
	return &server{
		app: app,
	}
}

func (s *server) Run(port int) {
	var mux http.Handler = s.routes()
	mux = corsMiddleware(mux)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	log.Printf("Starting server on :%d\n", port)
	log.Fatal(srv.ListenAndServe())
}
