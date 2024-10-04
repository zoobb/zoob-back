package server

import (
	"log"
	"net/http"
	"zoob-back/internal/db"
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
	todoDBCredentials := db.Credentials{
		User:     "zoob",
		Password: "1111",
		Name:     "todo",
		Host:     "localhost:9000",
	}
	db.Connect(todoDBCredentials)

	router := http.NewServeMux()
	router.HandleFunc("POST /ping", handler.Ping)
	router.HandleFunc("POST /auth/sign_up", handler.SignUp())
	router.HandleFunc("POST /auth/log_in", handler.LogIn())
	router.HandleFunc("GET /list", handler.GetAll())
	router.HandleFunc("DELETE /list", handler.DeleteAll())
	router.HandleFunc("POST /list", handler.AddToList())
	router.HandleFunc("GET /list/{id}", handler.ReadFromList())
	router.HandleFunc("PUT /list/{id}", handler.UpdateListItem())
	router.HandleFunc("DELETE /list/{id}", handler.DeleteListItem())

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
