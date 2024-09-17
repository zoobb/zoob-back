package server

import (
	"net/http"
	handlers "zoob-back/internal/delivery/http"
	"zoob-back/internal/models"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (s *Server) Run() error {
	todoList := &models.TodoList{
		Items: []string{},
	}

	router := http.NewServeMux()
	router.HandleFunc("POST /", handlers.Ping)
	router.HandleFunc("POST /list", handlers.AddToList(todoList))

	return http.ListenAndServe(s.addr, router)
}
