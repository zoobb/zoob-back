package http

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"zoob-back/internal/utils"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "There is an error occurred reading request: ", http.StatusInternalServerError)
	}
	log.Println("Request body: " + string(body))

	randomInt := utils.RandIntInRange(1, 100)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(strconv.Itoa(randomInt)))
	log.Println("Random number sent:", randomInt)
	if err != nil {
		return
	}
}
