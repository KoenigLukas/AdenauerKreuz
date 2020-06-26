package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/koeniglukas/config"
	"log"
)

var (
	host     = config.Get("DB_HOST")
	port     = config.Get("DB_PORT")
	user     = config.Get("DB_USER")
	password = config.Get("DB_PASSWORD")
	dbname   = config.Get("DB_NAME")
)

var (
	Con *sql.DB
)

func Init() {

	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	var err error
	Con, err = sql.Open("mysql", sqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = Con.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

