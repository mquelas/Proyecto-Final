package data_access

import (
	"database/sql"
	"log"
	//_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func dbConnect() (*sql.DB, error) {
	// conectarse a la DB
	var err error
	db, err = sql.Open("mysql", "arqsw:arqsw@tcp(127.0.0.1:3306)/arqsw")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	return db, err
}
