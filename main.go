package main

import (
	"log"
	"net/http"
	"pt-twwgt/resources"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/randomuser", resources.RandomUser)
	server := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}
	log.Fatal(server.ListenAndServe())
}
