package main

import (
	"./config"
	"./db"
	"./handler"
	"./middleware"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var Con sql.Conn

func main() {

	db.Init()
	defer db.Con.Close()

	router := mux.NewRouter()
	router.Use(middleware.LoggerMiddleware)

	router.HandleFunc("/api/v1/login", handler.LoginHandler).Methods("POST")
	router.HandleFunc("/api/v1/register", handler.RegisterHandler).Methods("POST")
	router.HandleFunc("/api/v1/check/userAvailable/{username}", handler.CheckUsernameHandler).Methods("GET")
	router.HandleFunc("/api/v1/check/emailAvailable/{email}", handler.CheckEmailHandler).Methods("GET")

	api := router.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.AuthMiddleware)

	api.HandleFunc("/test", handler.TestHandler).Methods("GET")

	api.HandleFunc("/list/create", handler.ListCreateHandler).Methods("POST")
	api.HandleFunc("/list/all", handler.ListGetAllHandler).Methods("POST")
	api.HandleFunc("/list/get/{id}", handler.ListGetHandler).Methods("GET")
	api.HandleFunc("/list/entry/add", handler.ListEntryAddHandler).Methods("POST")
	api.HandleFunc("/list/entry/edit", handler.ListEntryEditHandler).Methods("PUT")
	api.HandleFunc("/list/entry/delete", handler.ListEntryDeleteHandler).Methods("DELETE")
	api.HandleFunc("/user/change/password", handler.ChangePasswordHandler).Methods("PATCH")
	api.HandleFunc("/user/change/email", handler.ChangeEmailHandler).Methods("PATCH")
	api.HandleFunc("/user/delete", handler.DeleteHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+config.Get("SERVE_PORT"), router))

}


