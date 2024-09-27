package handler

import (
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"strconv"
)

type TodoList struct {
	Items []ListItem
}

type ListItem struct {
	ItemID  int    `json:"item_id"`
	Content string `json:"content"`
}

type ReqBody struct {
	UserData string `json:"user_data"`
}

/*func AddToList(list *TodoList, listItemID *int) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var reqBody ReqBody
		err := json.NewDecoder(req.Body).Decode(&reqBody)
		if err != nil {
			http.Error(rw, "Invalid request body", http.StatusBadRequest)
			return
		}
		log.Println("Request body:", reqBody)

		list.Items = append(list.Items, ListItem{
			ItemID:  *listItemID,
			Content: reqBody.UserData,
		})
		*listItemID += 1
		rw.WriteHeader(http.StatusOK)
	}
}*/

func AddToList(d *pgx.Conn) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

	}
}

func ReadFromList(list *TodoList) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		parsedID, err := strconv.Atoi(req.PathValue("id"))
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		if parsedID < 0 || parsedID >= len(list.Items) {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = rw.Write([]byte(list.Items[parsedID].Content))
		if err != nil {
			return
		}
		rw.WriteHeader(http.StatusOK)
	}
}

func UpdateListItem(list *TodoList) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		parsedID, err := strconv.Atoi(req.PathValue("id"))
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		if parsedID < 0 || parsedID >= len(list.Items) {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var reqBody ReqBody
		err = json.NewDecoder(req.Body).Decode(&reqBody)
		if err != nil {
			http.Error(rw, "Invalid request body", http.StatusBadRequest)
			return
		}
		log.Println("Request body:", reqBody)

		list.Items[parsedID].Content = reqBody.UserData
		rw.WriteHeader(http.StatusOK)
	}
}

func DeleteListItem(list *TodoList) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		parsedID, err := strconv.Atoi(req.PathValue("id"))
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		if parsedID < 0 || parsedID >= len(list.Items) {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		list.Items[parsedID].Content = ""
		rw.WriteHeader(http.StatusOK)
	}
}

func GetAll(list *TodoList) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(rw).Encode(list)
		if err != nil {
			return
		}
		rw.WriteHeader(http.StatusOK)
	}
}

func DeleteAll(list *TodoList, listItemID *int) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		list.Items = []ListItem{}
		*listItemID = 0

		rw.WriteHeader(http.StatusOK)
	}
}
