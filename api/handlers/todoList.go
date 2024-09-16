package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type AddToListReqBody struct {
	ListItem string `json:"listItem"`
}

type TodoList struct {
	Items []string
}

func (l *TodoList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var reqBody AddToListReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Println("Request body:", reqBody)

	l.Items = append(l.Items, reqBody.ListItem)
	log.Println(l)
	w.WriteHeader(http.StatusOK)
}

/*
func AddToList(l []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		var reqBody AddToListReqBody
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		log.Println("Request body:", reqBody)

		l = append(l, reqBody.ListItem)
		log.Println(l)
		w.WriteHeader(http.StatusOK)
	}
}
*/
