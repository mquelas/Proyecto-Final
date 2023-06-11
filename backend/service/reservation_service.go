package service

import (
	. "backend/data_access"
	. "backend/model"
	"database/sql"
	"fmt"
)

func GetReservations() ([]Reservation, error) {
	var err error
	var db *sql.DB

	db, err = DataConnect()
	if err != nil {
		return nil, fmt.Errorf("getReservations %q", err)
	}
	defer db.Close()

	var reservations []Reservation

	rows, err := db.Query(`SELECT * FROM reservations`)
	if err != nil {
		return nil, fmt.Errorf("getReservations %q", err)
	}
	defer rows.Close()

	for rows.Next() {
		var reser Reservation
		if err := rows.Scan(&reser.ID, &reser.IdHotel, &reser.CheckIn, &reser.CheckOut, &reser.IsConfirmed); err != nil {
			return nil, fmt.Errorf("getReservations %q", err)
		}
		reservations = append(reservations, reser)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getReservations %q", err)
	}

	return reservations, nil
}

func GetReservationById(id string) (Reservation, error) {

	var reservation Reservation
	var err error
	var db *sql.DB

	db, err = DataConnect()

	if err != nil {
		return reservation, fmt.Errorf("getReservationById %q", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM reservations WHERE Id_reservations = ?", id)
	if err != nil {
		return reservation, fmt.Errorf("getReservationById %q", err)
	}
	defer rows.Close()

	for rows.Next() {
		var reser Reservation
		if err := rows.Scan(&reser.ID, &reser.IdHotel, &reser.CheckIn, &reser.CheckOut, &reser.IsConfirmed); err != nil {
			return reservation, fmt.Errorf("getReservationById %q", err)
		}
		reservation = reser
	}

	if err := rows.Err(); err != nil {
		return reservation, fmt.Errorf("getReservationById %q", err)
	}

	return reservation, nil
}
