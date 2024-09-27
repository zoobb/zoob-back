package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"zoob-back/internal/db"
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

func AddToList() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var reqBody ReqBody
		err := json.NewDecoder(req.Body).Decode(&reqBody)
		if err != nil {
			http.Error(rw, "Invalid request body", http.StatusBadRequest)
			return
		}
		log.Println("Request body:", reqBody)

		err = db.AddToList(reqBody.UserData)
		if err != nil {
			log.Println(err)
			return
		}

		rw.WriteHeader(http.StatusOK)
	}
}

/*func ReadFromList(list *TodoList) http.HandlerFunc {
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
}*/

func ReadFromList() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		parsedID, err := strconv.Atoi(req.PathValue("id"))
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		var selection string
		selection, err = db.ReadFromList(parsedID)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = rw.Write([]byte(selection))
		if err != nil {
			log.Println(err)
			return
		}

		rw.WriteHeader(http.StatusOK)
	}
}

/*func UpdateListItem(list *TodoList) http.HandlerFunc {
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
}*/

func UpdateListItem() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		parsedID, err := strconv.Atoi(req.PathValue("id"))
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}

		var reqBody ReqBody
		err = json.NewDecoder(req.Body).Decode(&reqBody)
		if err != nil {
			http.Error(rw, "Invalid request body", http.StatusBadRequest)
			log.Println(err)
			return
		}
		log.Println("Request body:", reqBody)
		err = db.UpdateListItem(parsedID, reqBody.UserData)
		if err != nil {
			log.Println(err)
			return
		}
		rw.WriteHeader(http.StatusOK)
	}
}

/*func DeleteListItem(list *TodoList) http.HandlerFunc {
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
}*/

func DeleteListItem() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		parsedID, err := strconv.Atoi(req.PathValue("id"))
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}

		err = db.DeleteListItem(parsedID)
		if err != nil {
			log.Println(err)
			return
		}
		rw.WriteHeader(http.StatusOK)
	}
}

/*func GetAll(list *TodoList) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(rw).Encode(list)
		if err != nil {
			log.Println(err)
			return
		}
		rw.WriteHeader(http.StatusOK)
	}
}*/

func GetAll() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		list, err := db.GetAll()
		if err != nil {
			log.Println(err)
			return
		}
		err = json.NewEncoder(rw).Encode(list)

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
