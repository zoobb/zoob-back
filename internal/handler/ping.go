package handler

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"zoob-back/internal/utils"
)

func Ping(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "There is an error occurred reading request: ", http.StatusInternalServerError)
	}
	log.Println("Request reqBody: " + string(reqBody))

	randomInt := utils.RandIntInRange(1, 100)

	_, err = rw.Write([]byte(strconv.Itoa(randomInt)))
	log.Println("Random number sent:", randomInt)
	if err != nil {
		return
	}
	rw.WriteHeader(http.StatusOK)
}
