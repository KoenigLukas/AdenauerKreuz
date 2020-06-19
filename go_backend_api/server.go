package main

import (
	"github.com/gorilla/mux"
	"net/http"
)


func main() {
	router := mux.NewRouter()
	http.ListenAndServe(":8000", router)
	router.HandleFunc("/v1/login", loginHandler).Methods("POST")
}
