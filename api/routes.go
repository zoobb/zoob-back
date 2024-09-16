package api

import (
	"net/http"
	"zoob-back/api/handlers"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) RouteReg(router *http.ServeMux) {
	todoList := &handlers.TodoList{
		Items: []string{},
	}

	router.HandleFunc("POST /", handlers.Ping)
	router.Handle("POST /list", todoList)
}
