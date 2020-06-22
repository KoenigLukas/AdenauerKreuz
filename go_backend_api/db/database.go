package db

import (
	"../config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	host     = config.Get("DB_HOST")
	port     = config.Get("DB_PORT")
	user     = config.Get("DB_USER")
	password = config.Get("DB_PASSWORD")
	dbname   = config.Get("DB_NAME")
)

var Db *sql.DB

func Init() {

	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	db, err := sql.Open("mysql", sqlInfo)
	if err != nil {
		log.Fatal(err)
	}
}
