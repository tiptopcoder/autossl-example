//main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", DoHealthCheck).Methods("GET")
	log.Println("Golang server starts at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
func DoHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, i'm a golang microservice")
	w.WriteHeader(http.StatusAccepted) //RETURN HTTP CODE 202
}
