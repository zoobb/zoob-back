package server

import (
	"log"
	"net/http"
	"zoob-back/internal/handler"
)

type Server struct {
	addr string
}

func New(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

type NewType []bool

func (n *NewType) hmmm() {
	log.Println("hmmm")
}

func (s *Server) Run() error {
	todoList := &handler.TodoList{
		Items: []handler.ListItem{},
	}
	listItemID := 0

	router := http.NewServeMux()
	router.HandleFunc("POST /", handler.Ping)
	router.HandleFunc("POST /add", handler.AddToList(todoList, &listItemID))
	router.HandleFunc("GET /read/{id}", handler.ReadFromList(todoList))
	router.HandleFunc("PUT /update/{id}", handler.UpdateListItem(todoList))
	router.HandleFunc("DELETE /delete/{id}", handler.DeleteListItem(todoList))

	return http.ListenAndServe(s.addr, router)
}
