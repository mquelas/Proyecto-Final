package service

import (
	. "backend/data_access"
	. "backend/model"
	"database/sql"
	"fmt"
)

func GetReservations([]Reservation, error) {

	var err error
	var db *sql.DB

	db, err = DataConnect()
	if err != nil {

		return nil fmt.Errorf("getReservations %q", err)
	}

	defer db.Close()

	var reservations []Reservation

	rows, err := db.Query(`SELECT * FROM reservations`)
	if err != nil {
		return nil fmt.Errorf("getReservations %q", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.

	for rows.Next() {

		var reser Reservation
		if err := rows.Scan(&reser.ID, &reser.IdHotel, &reser.CheckIn, &reser.CheckOut, &reser.IsConfirmed); err != nil {
			return nil fmt.Errorf("getReservations %q", err)
		}
		reservations = append(reservations, reser)
	}
	if err := rows.Err(); err != nil {
		return nil
	}

	return reservations
}

func GetReservationById(id string) []Reservation {

	var reservations []Reservation
	var err error
	var db *sql.DB

	rows, err := db.Query("SELECT * FROM reservations WHERE Id = ?", id)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {

		var reser Reservation
		if err := rows.Scan(&reser.ID, &reser.IdHotel, &reser.CheckIn, &reser.CheckOut, &reser.IsConfirmed); err != nil {
			return nil
		}
		reservations = append(reservations, reser)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return reservations
}
