package api

import "net/http"

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (s *Server) Run() error {
	router := http.NewServeMux()
	handler := NewHandler()
	handler.RouteReg(router)
	return http.ListenAndServe(s.addr, router)
}
