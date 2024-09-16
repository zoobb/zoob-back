package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type AddToListReqBody struct {
	ListItem string `json:"listItem"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var todoList []string
	port := 8247
	/*var serverURL = url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("localhost:%d", port),
	}*/

	http.HandleFunc("POST /", handlePingReq)
	http.HandleFunc("POST /list", handleAddToListReq(todoList))
	// I KNOW
	log.Println("Server started on port ", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("There's an error occurred running server: ", err)
		return
	}
}

func handlePingReq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "There is an error occurred reading request: ", http.StatusInternalServerError)
	}
	log.Println("Request body: " + string(body))

	randomInt := randIntInRange(1, 100)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(strconv.Itoa(randomInt)))
	log.Println("Random number sent:", randomInt)
	if err != nil {
		return
	}
}

func handleAddToListReq(l []string) http.HandlerFunc {
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

func randIntInRange(min int, max int) int {
	return rand.Intn(max-min) + min
}
