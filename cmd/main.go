package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"zoob-back/internal/server"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	port := 8247
	server := server.NewServer(fmt.Sprintf(":%d", port))
	// I KNOW
	log.Println("Server started on port", port)
	err := server.Run()
	if err != nil {
		log.Fatal("There's an error occurred running server:", err)
		return
	}
}
