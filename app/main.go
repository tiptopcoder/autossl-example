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
	router.HandleFunc("/domain", CheckDomain).Methods("GET")
	log.Println("Golang server starts at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
func DoHealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host)
	fmt.Fprintf(w, "Hello, i'm a golang microservice "+r.Host)
	w.WriteHeader(http.StatusAccepted) //RETURN HTTP CODE 202
}
func CheckDomain(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("s")
	log.Println(domain)

	if domain == "cuiseene.com" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
