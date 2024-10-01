package handler

import (
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"strconv"
	"zoob-back/internal/db"
)

type ReqBody struct {
	UserData string `json:"user_data"`
}

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
		if errors.Is(err, pgx.ErrNoRows) {
			rw.WriteHeader(http.StatusNotFound)
			log.Println(err)
			return
		} else if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
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
		log.Println("Request body:", reqBody, parsedID)
		err = db.UpdateListItem(parsedID, reqBody.UserData)
		if errors.Is(err, pgx.ErrNoRows) {
			rw.WriteHeader(http.StatusNotFound)
			log.Println(err)
			return
		} else if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		rw.WriteHeader(http.StatusOK)
	}
}
func DeleteListItem() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		parsedID, err := strconv.Atoi(req.PathValue("id"))
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}

		err = db.DeleteListItem(parsedID)
		if errors.Is(err, pgx.ErrNoRows) {
			rw.WriteHeader(http.StatusNotFound)
			log.Println(err)
			return
		} else if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		rw.WriteHeader(http.StatusOK)
	}
}
func GetAll() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		list, err := db.GetAll()
		if err != nil {
			log.Println(err)
			return
		}
		err = json.NewEncoder(rw).Encode(list)
	}
}
func DeleteAll() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		err := db.DeleteAll()
		if err != nil {
			return
		}

		rw.WriteHeader(http.StatusOK)
	}
}
