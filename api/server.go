package api

import "net/http"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	return http.ListenAndServe(":9091", nil)
}
