package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Hotel struct {
	ID          string  `json:"id"`
	Name        string  `json: name`
	Description string  `json: description`
	Price       float64 `json: price`
}

func main() {
	// propiedades de la conexion a la DB
	/*
		cfg := mysql.Config{
			//User:   os.Getenv("DBUSER"),
			User: "arqsw",
			//Passwd: os.Getenv("DBPASS"),
			Passwd: "arqsw",
			Net:    "tcp",
			Addr:   "127.0.0.1:3306",
			DBName: "arqsw",
		}
	*/
	// conectarse a la DB
	var err error
	//db, err = sql.Open("mysql", cfg.FormatDSN())
	db, err = sql.Open("mysql", "arqsw:arqsw@tcp(127.0.0.1:3306)/arqsw")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	fmt.Println(hotelsById("1"))
}

// hotelsById queries for hotel that have the specified Id.

func hotelsById(name string) ([]Hotel, error) {

	// An hotel slice to hold data from returned rows.

	var hotels []Hotel

	rows, err := db.Query("SELECT * FROM hotel WHERE Id = ?", name)
	if err != nil {
		return nil, fmt.Errorf("hotelsById %q: %v", name, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.

	for rows.Next() {

		var hot Hotel
		if err := rows.Scan(&hot.ID, &hot.Name, &hot.Description, &hot.Price); err != nil {
			return nil, fmt.Errorf("hotelsById %q: %v", name, err)
		}
		hotels = append(hotels, hot)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("hotelsById %q: %v", name, err)
	}
	return hotels, nil
}
