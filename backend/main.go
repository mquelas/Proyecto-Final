package main

import (
	//"net/http"

	. "backend/model"
	. "backend/service"
	"net/http"

	//. "backend/data_access"
	//"fmt"
	"github.com/gin-gonic/gin"
)

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

// insertar neuvos hoteles

// ---------------------------------------------------------------------------------------------
func postHotels(context *gin.Context) {

	var newHotel Hotel
	if err := context.BindJSON(&newHotel); err != nil {
		return
	}

	context.IndentedJSON(http.StatusCreated, newHotel)
}

// ---------------------------------------------------------------------------------------------
func getHotels(context *gin.Context) {
	var hotels, err = GetHotels()
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}
	context.IndentedJSON(http.StatusOK, hotels)
}

func getHotelById(context *gin.Context) {
	id := context.Param("id")

	var hotel, err = GetHotelById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, hotel)
}

func getReservations(context *gin.Context) {

	reservations, err := GetReservations()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, reservations)
}

func getReservationById(context *gin.Context) {

	var reservation Reservation
	var err error
	id := context.Param("id")

	reservation, err = GetReservationById(id)
	if err != nil {

		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, reservation)
}

// ----------------------------------------------------------------------------------------------------------------------
func postReservations(context *gin.Context) {

	var newReservation Reservation
	var err error

	if err = context.BindJSON(&newReservation); err != nil {
		return
	}

	newReservation, err = CreateReservation(newReservation)

	context.IndentedJSON(http.StatusCreated, newReservation)
}

//----------------------------------------------------------------------------------------------------------------------

//confirma la reserva

func confirmReservation(reservation *Reservation) {

	reservation.IsConfirmed = true
}
