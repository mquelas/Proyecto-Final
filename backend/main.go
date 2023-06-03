package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "backend/model"
	. "backend/service"
)

// estructura reservas
type reservation struct {
	ID          string `json:"id"`
	IsConfirmed bool   `json:"isConfirmed"`
	IdHotel     string `json:"idHotel"`
	From        string `json:"from"`
	To          string `json:"to"`
}

//confirma la reserva

func confirmReservation(reservation *reservation) {

	reservation.IsConfirmed = true
}

//crea resevas

var reservations = []reservation{

	{ID: "1", IsConfirmed: false, IdHotel: "1", From: "2023-05-23", To: "2023-06-03"},
	{ID: "2", IsConfirmed: false, IdHotel: "2", From: "2024-05-23", To: "2025-06-03"},
	{ID: "3", IsConfirmed: false, IdHotel: "3", From: "2020-05-23", To: "2020-06-03"},
}

/*
	GET --- listar hoteles
	agregar un hotel
	ver detalles de un hotel

	GET --- listar reservas
	confirmar una reserva ----------- FALTA VALIDACION

*/

func main() {

	router := gin.Default()
	router.GET("/hotels", getHotels)
	router.GET("/hotels/:id", getHotelById)
	router.POST("/hotels", postHotels)
	router.GET("/reservations", getReservations)
	router.GET("/reservations/:id", getReservationById)
	router.POST("/reservations", postReservations)

	router.Run("localhost:8080")
}

func getHotels(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, GetHotels())
}

func postHotels(context *gin.Context) {
	var newHotel Hotel
	if err := context.BindJSON(&newHotel); err != nil {
		return
	}

	//hotels = append(hotels, newHotel)
	context.IndentedJSON(http.StatusCreated, newHotel)
}

func getHotelById(context *gin.Context) {
	//id := context.Param("id")
	/*
		for _, hotel := range hotels {
			if hotel.ID == id {
				context.IndentedJSON(http.StatusOK, hotel)
				return
			}
		}
	*/context.IndentedJSON(http.StatusNotFound, gin.H{"message": "hotel no encontrado"})
}

func getReservations(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, reservations)
}

func getReservationById(context *gin.Context) {
	id := context.Param("id")

	for _, reservation := range reservations {
		if reservation.ID == id {
			context.IndentedJSON(http.StatusOK, reservation)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "reserva no encontrada"})
}

func postReservations(context *gin.Context) {
	var newReservation reservation
	if err := context.BindJSON(&newReservation); err != nil {
		return
	}

	reservations = append(reservations, newReservation)
	context.IndentedJSON(http.StatusCreated, newReservation)
}
