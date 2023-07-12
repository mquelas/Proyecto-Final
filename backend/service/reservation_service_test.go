package service

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"
)

func TestGetReservations_Success(t *testing.T) {

	var db *sql.DB

	expectedReservations := []Reservation{

		{ID: 1, CheckIn: "2023-01-01", CheckOut: "2023-01-03", IdHotel: 1, EMail: "prueba@gmail.com"},
		{ID: 2, CheckIn: "2023-07-04", CheckOut: "2023-07-07", IdHotel: 2, EMail: "pepcito@gmail.com"},
	}

	_ = expectedReservations

	//nose como hacer si no mando un query aca

	db.Query = func(query string, args ...interface{}) (*sql.Rows, error) {

		rows := &mockSQLRows{ //esto de mock lo lei en googlpero no lo enti muy bien

			data: [][]interface{}{

				{1, "2023-01-01", "2023-01-03", 1, "prueba@gmail.com"},
				{2, "2023-07-04", "2023-07-07", 2, "pepcito@gmail.com"},
			},
		}
		return rows, nil
	}

	reservations, err := GetReservations()

	_ = reservations //nose que tan bien es hacer eso en todo

	if err != nil {

		t.Fatalf("Error al invocar GetReservations: %v", err)
	}
	//falta ver si hace falta otro if

}

func TestGetReservations_Failure(t *testing.T) {

	var db *sql.DB
	//mimo error de arriba, no me deja usar query (que saqe de google que se podia hacer asi)
	db.Query = func(query string, args ...interface{}) (*sql.Rows, error) {

		return nil, errors.New("Error en la ejecución de la consulta")
	}

	_, err := GetReservations()

	if err == nil {

		t.Fatalf("GetReservations no falla cuando hay un error en la ejecución de la consulta")
	}
}

func TestGetReservationById_Success(t *testing.T) {

	var db *sql.DB
	id := int64(1)
	//reservation no esta definida
	expectedReservation := Reservation{
		ID:       1,
		CheckIn:  "2023-01-01",
		CheckOut: "2023-01-03",
		IdHotel:  1,
		EMail:    "prueba@gmail.com",
	}
	//mismo problema de query
	db.Query = func(query string, args ...interface{}) (*sql.Rows, error) {

		//mismo problema de arriba
		rows := &mockSQLRows{

			data: [][]interface{}{

				{1, "2023-01-01", "2023-01-03", 1, "prueba@gmail.com"},
			},
		}
		return rows, nil
	}

	reservation, err := GetReservationById(id)

	if err != nil {

		t.Fatalf("Error al invocar GetReservationById: %v", err)
	}

	if !compareReservations(reservation, expectedReservation) {

		t.Fatalf("La reserva retornada no coincide con la reserva esperada")
	}
}

func TestGetReservationById_Failure(t *testing.T) {

	var db *sql.DB

	id := int64(1)
	//la solucion que me da es hacer (sql.DB).Query pero tampoco funciona
	db.Query = func(query string, args ...interface{}) (*sql.Rows, error) {

		return nil, errors.New("Error en la ejecución de la consulta")
	}

	_, err := GetReservationById(id)

	if err == nil {

		t.Fatalf("GetReservationById no falla cuando hay un error en la ejecución de la consulta")
	}
}

// no entendi, sacado de videoito de youtube
func compareReservations(reservation1, reservation2 Reservation) bool {

	return reservation1.ID == reservation2.ID &&

		reservation1.CheckIn == reservation2.CheckIn &&
		reservation1.CheckOut == reservation2.CheckOut &&
		reservation1.IdHotel == reservation2.IdHotel &&
		reservation1.EMail == reservation2.EMail
}

func TestCreateReservation_Success(t *testing.T) {

	reservation := Reservation{
		CheckIn:  time.Now().AddDate(0, 0, 1),
		CheckOut: time.Now().AddDate(0, 0, 3),
		IdHotel:  1,
		EMail:    "prueba@gmail.com",
	}

	validateAvailability = func(reservation Reservation) bool {

		return true
	}

	dbExecContext = func(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {

		return &mockSQLResult{lastInsertID: 1}, nil
	}

	insertOccupancy = func(db *sql.DB, reservationID int64, date time.Time) error {

		return nil
	}

	newReservation, err := CreateReservation(reservation)

	if err != nil {
		t.Fatalf("Error al invocar CreateReservation: %v", err)
	}

	if newReservation.ID != 1 {
		t.Fatalf("El ID de la nueva reserva no coincide con el esperado")
	}
}

func TestCreateReservation_AvailabilityError(t *testing.T) {

	reservation := Reservation{
		CheckIn:  time.Now().AddDate(0, 0, 1),
		CheckOut: time.Now().AddDate(0, 0, 3),
		IdHotel:  1,
		EMail:    "prueba@gmail.com",
	}

	validateAvailability = func(reservation Reservation) bool {

		return false
	}

	_, err := CreateReservation(reservation)

	if err == nil {

		t.Fatalf("CreateReservation no falla cuando no hay disponibilidad")
	}
}

func TestCreateReservation_DatabaseError(t *testing.T) {

	reservation := Reservation{
		CheckIn:  time.Now().AddDate(0, 0, 1),
		CheckOut: time.Now().AddDate(0, 0, 3),
		IdHotel:  1,
		EMail:    "prueba@gmail.com",
	}

	validateAvailability = func(reservation Reservation) bool {

		return true
	}

	dbExecContext = func(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {

		return nil, errors.New("Error en la ejecución de la consulta")
	}

	_, err := CreateReservation(reservation)

	if err == nil {

		t.Fatalf("CreateReservation no falla cuando hay un error en la ejecución de la consulta")
	}
}

func TestCreateReservation_OccupancyError(t *testing.T) {

	reservation := Reservation{

		CheckIn:  time.Now().AddDate(0, 0, 1),
		CheckOut: time.Now().AddDate(0, 0, 3),
		IdHotel:  1,
		EMail:    "prueba@gmail.com",
	}

	validateAvailability = func(reservation Reservation) bool {

		return true
	}

	dbExecContext = func(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
		// Mock the LastInsertId to return 1
		return &mockSQLResult{lastInsertID: 1}, nil
	}

	insertOccupancy = func(db *sql.DB, reservationID int64, date time.Time) error {
		return errors.New("Error en la inserción de ocupación")
	}

	// Act
	_, err := CreateReservation(reservation)

	if err == nil {

		t.Fatalf("CreateReservation no falla cuando hay un error en la inserción de ocupación")
	}
}
