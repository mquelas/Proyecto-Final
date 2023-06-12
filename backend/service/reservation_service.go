package service

import (
	. "backend/data_access"
	. "backend/model"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
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

	rows, err := db.Query("SELECT * FROM reservation WHERE Id_reservation = ?", id)

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

func CreateReservation(reservation Reservation) (Reservation, error) {

	var err error
	var db *sql.DB
	//var reservation Reservation

	db, err = DataConnect()

	if err != nil {

		return reservation, fmt.Errorf("createReservation %q", err)
	}

	defer db.Close()

	insertResult, err := db.ExecContext(

		context.Background(),
		"INSERT INTO reservation (id_reservation, checkin, checkout, id_hotel, IsConfirmed) VALUES (?, ?, ?, ?, ?)",
		reservation.ID, reservation.CheckIn, reservation.CheckOut, reservation.IdHotel, reservation.IsConfirmed,
	)

	if err != nil {

		log.Fatalf("impossible insert reservation: %s", err)
	}

	id, err := insertResult.LastInsertId()

	if err != nil {

		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}

	log.Printf("inserted id: %d", id)

	if err != nil {

		return reservation, fmt.Errorf("createReservation %q", err)
	}

	return reservation, nil
}

func GetReservationByHotelId(id string) (Reservation, error) {

	var reservation Reservation
	var err error
	var db *sql.DB

	db, err = DataConnect()

	if err != nil {

		return reservation, fmt.Errorf("getReservationById %q", err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM reservation WHERE Id_hotel = ?", id)

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

func getReservationByEmail(id string) (Reservation, error) {

	var reservation Reservation
	var err error
	var db *sql.DB

	db, err = DataConnect()

	if err != nil {

		return reservation, fmt.Errorf("getReservationById %q", err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM reservation WHERE email = ?", id)

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
func InsertReservationIntoOccupancy(reservationID int, roomID int, startDate time.Time, endDate time.Time) error {

	var err error
	var db *sql.DB

	db, err = DataConnect()

	if err != nil {

		return fmt.Errorf("insertReservation: %q", err)
	}
	defer db.Close()

	// Insertar la reserva en la tabla "Reservation"

	insertResult, err := db.ExecContext(

		context.Background(),
		"INSERT INTO reservation (id_reservation, checkin, checkout, id_hotel, IsConfirmed) VALUES (?, ?, ?, ?, ?)",
		reservationID, startDate, endDate, roomID, false,
	)
	if err != nil {

		return fmt.Errorf("impossible to insert reservation: %s", err)
	}

	// Obtener el ID de la reserva insertada

	id, err := insertResult.LastInsertId() //-----------------------VER BIEN--------------------------------------

	if err != nil {

		return fmt.Errorf("No se pudo traer el LastInsertId: %s", err)
	}

	log.Printf("inserted ID: %d", id)

	// Insertar la reserva en la tabla "Occupancy" para cada día entre checkin y checkout

	for date := startDate; date.Before(endDate); date = date.AddDate(0, 0, 1) {

		//----------------------------------FALTA COMPLETAR--------------------------------------------
	}

	return nil
}
