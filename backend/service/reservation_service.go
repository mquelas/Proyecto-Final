package service

import (
	. "backend/data_access"
	. "backend/model"
	"context"
	"database/sql"
	"fmt"
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

	rows, err := db.Query(`SELECT * FROM reservation`)

	if err != nil {
		return nil, fmt.Errorf("getReservations %q", err)
	}

	defer rows.Close()

	for rows.Next() {
		var reser Reservation

		if err := rows.Scan(&reser.ID, &reser.CheckIn, &reser.CheckOut, &reser.IdHotel, &reser.EMail); err != nil {
			return nil, fmt.Errorf("getReservations %q", err)
		}

		reservations = append(reservations, reser)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getReservations %q", err)
	}

	return reservations, nil
}

func GetReservationById(id int64) (*Reservation, error) {

	var reservation Reservation
	var err error
	var db *sql.DB

	db, err = DataConnect()

	if err != nil {
		return nil, fmt.Errorf("getReservationById %q", err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM reservation WHERE Id_reservation = ?", id)

	if err != nil {
		return nil, fmt.Errorf("getReservationById %q", err)
	}

	defer rows.Close()

	var reser Reservation

	for rows.Next() {

		if err := rows.Scan(&reser.ID, &reser.CheckIn, &reser.CheckOut, &reser.IdHotel, &reser.EMail); err != nil {
			return nil, fmt.Errorf("getReservationById %q", err)
		}

		reservation = reser
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getReservationById %q", err)
	}
	if reser.ID == int64(0) {
		return nil, fmt.Errorf("reservation no existe")
	}

	return &reservation, nil
}

func CreateReservation(reservation Reservation) (*Reservation, error) {

	if !validateAvailability(reservation) {
		return nil, fmt.Errorf("No hay disponibilidad")
	}

	var err error
	var db *sql.DB

	db, err = DataConnect()
	if err != nil {
		return nil, fmt.Errorf("createReservation %q", err)
	}

	defer db.Close()

	insertResult, err := db.ExecContext(
		context.Background(),
		"INSERT INTO reservation (checkin, checkout, id_hotel, Email) VALUES (?, ?, ?, ?)",
		reservation.CheckIn, reservation.CheckOut, reservation.IdHotel, reservation.EMail,
	)
	if err != nil {
		return nil, fmt.Errorf("no se pudo insertar la reservacion %q", err)
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("no se pudo obtener el ID insertado %q", err)
	}

	var newReservation *Reservation = &reservation

	if newReservation == nil {

		return nil, fmt.Errorf("La reservacion es nula %q", err)
	}
	newReservation.ID = id

	diffDates := reservation.CheckOut.Sub(reservation.CheckIn).Hours() / 24

	for i := 0; i < int(diffDates); i++ {
		err := insertOccupancy(
			db,
			id,
			reservation.CheckIn.AddDate(0, 0, i))

		if err != nil {
			return nil, fmt.Errorf("createReservation %q", err)
		}
	}

	return newReservation, nil
}

func validateAvailability(reservation Reservation) bool {
	hotel, err := GetHotelById(reservation.IdHotel)
	if err != nil {
		return false
	}

	var qtyRooms = hotel.Rooms
	reservedRooms, err := GetOccupancyByDate(reservation.IdHotel, reservation.CheckIn, reservation.CheckOut)
	if err != nil || reservedRooms < 0 {
		return false
	}

	if qtyRooms > reservedRooms {
		return true
	}

	return false
}

func insertOccupancy(db *sql.DB, reservationID int64, date time.Time) error {
	var err error

	insertResult, err := db.ExecContext(
		context.Background(),
		"INSERT INTO occupancy (id_reservation, date) VALUES (?, ?)",
		reservationID, date,
	)
	if err != nil {
		return fmt.Errorf("impossible to insert occupancy: %s", err)
	}

	id, err := insertResult.LastInsertId()
	if err != nil || id < 1 {
		return fmt.Errorf("No se pudo traer el LastInsertId: %s", err)
	}

	return nil
}

func GetReservationByHotelId(id int64) ([]Reservation, error) {
	var reservations []Reservation
	var err error
	var db *sql.DB

	db, err = DataConnect()

	if err != nil {
		return nil, fmt.Errorf("getReservationByHotelId %q", err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM reservation WHERE Id_hotel = ?", id)

	if err != nil {
		return nil, fmt.Errorf("getReservationByHotelId %q", err)
	}

	defer rows.Close()

	for rows.Next() {

		var reser Reservation

		if err := rows.Scan(&reser.ID, &reser.CheckIn, &reser.CheckOut, &reser.IdHotel, &reser.EMail); err != nil {
			return nil, fmt.Errorf("getReservationById %q", err)
		}

		reservations = append(reservations, reser)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getReservationById %q", err)
	}

	return reservations, nil
}

func GetReservationByEmail(email string) ([]Reservation, error) {

	var reservations []Reservation
	var err error
	var db *sql.DB

	db, err = DataConnect()

	if err != nil {
		return nil, fmt.Errorf("getReservationById %q", err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM reservation WHERE email = ?", email)

	if err != nil {
		return nil, fmt.Errorf("getReservationById %q", err)
	}
	defer rows.Close()

	for rows.Next() {
		var reser Reservation

		if err := rows.Scan(&reser.ID, &reser.CheckIn, &reser.CheckOut, &reser.IdHotel, &reser.EMail); err != nil {
			return nil, fmt.Errorf("getReservationById %q", err)
		}
		reservations = append(reservations, reser)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getReservationById %q", err)
	}

	return reservations, nil
}
