package http

import (
	"encoding/json"
	"log"
	"net/http"
	"zoob-back/internal/models"
)

func AddToList(l *models.TodoList) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		var reqBody models.AddToListReqBody
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
}
