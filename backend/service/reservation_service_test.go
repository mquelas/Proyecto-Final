package service

import (
	. "backend/model"
	"testing"
	"time"
)

func TestGetReservations_Success(t *testing.T) {
	// se asume que existen registros en Reservation
	reservations, err := GetReservations()

	if err != nil {
		t.Fatalf("Error al invocar GetReservations: %v", err)
	}
	if len(reservations) == 0 {
		t.Fatalf("No se encontraron reservations %v", err)
	}

}

func TestGetReservationById_Success(t *testing.T) {
	// se asume que existe un registro en Reservation con ID = 16
	var idReservation = int64(16)
	reservation, err := GetReservationById(idReservation)

	if err != nil {
		t.Fatalf("Error al invocar GetReservationById: %v", err)
	}

	if reservation.ID != idReservation {
		t.Fatalf("La reserva retornada no coincide con la reserva esperada")
	}
}

func TestGetReservationById_Failure(t *testing.T) {
	// se asume que NO existe un registro en Reservation con ID = 1
	var idReservation = int64(1)
	reservation, err := GetReservationById(idReservation)

	if err == nil {
		t.Fatalf("GetReservationById no falla cuando hay un error en la ejecución de la consulta")
	}
	if reservation != nil && reservation.ID != 0 {
		t.Fatalf("GetReservationById devuelve una Reservation que no deberia devolver")
	}
}

func TestCreateReservation_Success(t *testing.T) {

	reservation := Reservation{
		CheckIn:  time.Now().AddDate(0, 0, 1),
		CheckOut: time.Now().AddDate(0, 0, 3),
		IdHotel:  1,
		EMail:    "prueba@gmail.com",
	}

	newReservation, err := CreateReservation(reservation)

	if err != nil {
		t.Fatalf("Error al invocar CreateReservation: %v", err)
	}

	if newReservation.ID == 0 {
		t.Fatalf("El ID de la nueva reserva no es valido")
	}
}

func TestCreateReservation_AvailabilityError(t *testing.T) {

	//TODO ojo que para que funcione el test hay que tener datos en la db cargados
	reservation := Reservation{
		CheckIn:  time.Now().AddDate(0, 0, 1),
		CheckOut: time.Now().AddDate(0, 0, 3),
		IdHotel:  1,
		EMail:    "prueba@gmail.com",
	}

	_, err := CreateReservation(reservation)

	if err == nil {
		t.Fatalf("CreateReservation no falla cuando no hay disponibilidad")
	} else {
		if err.Error() != "No hay disponibilidad" {
			t.Fatalf("CreateReservation no indica que no hay disponibilidad")
		}
	}
}
