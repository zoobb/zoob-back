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
	router.HandleFunc("POST /ping", handler.Ping)
	router.HandleFunc("GET /list", handler.GetAll(todoList))
	router.HandleFunc("POST /list", handler.AddToList(todoList, &listItemID))
	router.HandleFunc("GET /list/{id}", handler.ReadFromList(todoList))
	router.HandleFunc("PUT /list/{id}", handler.UpdateListItem(todoList))
	router.HandleFunc("DELETE /list/{id}", handler.DeleteListItem(todoList))

	return http.ListenAndServe(s.addr, withCORS(router))
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if req.Method == "OPTIONS" {
			rw.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(rw, req)
	})
}
