package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("ET")
	yRouter.HandleFunc()
log.Fatal(http.ListenAndServe(":8081", myRouter))
}

unc main() {
	fmt.Println("Go RM sample")

	handleRequests()
}
