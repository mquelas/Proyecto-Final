package service

import (
	. "backend/data_access"
	. "backend/model"
	"context"
	"database/sql"
	"fmt"
	"log"
)

func CreateUser(user User) (User, error) {

	var err error
	var db *sql.DB
	//var reservation Reservation

	db, err = DataConnect()

	if err != nil {

		return user, fmt.Errorf("createUser %q", err)
	}

	defer db.Close()

	insertResult, err := db.ExecContext(
		context.Background(),
		"INSERT INTO user (email, name, lastname, password, admin) VALUES (?, ?, ?, ?, ?)",
		user.EMail, user.Name, user.LastName, user.Password, user.Admin,
	)

	if err != nil {
		log.Fatalf("impossible insert user: %s", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}
	log.Printf("inserted id: %d", id)

	if err != nil {
		return user, fmt.Errorf("createUser %q", err)
	}

	return user, nil
}
