package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicboul/shortd/internal/reader"
	"github.com/nicboul/shortd/internal/writer"
)

type Server struct {
	Mux    *mux.Router
	Listen string
}

func NewServer(listen string) *Server {
	server := &Server{
		Listen: listen,
		Mux:    mux.NewRouter(),
	}

	writer := writer.NewWriter()
	reader := reader.NewReader()

	server.Mux.Methods("POST").PathPrefix("/").Handler(writer)
	server.Mux.Methods("GET").PathPrefix("/").Handler(reader)

	return server
}

func (s *Server) Serve() {

	httpServer := &http.Server{
		Addr:    s.Listen,
		Handler: s.Mux,
	}
	httpServer.ListenAndServe()
}
