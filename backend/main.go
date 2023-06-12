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
	router.POST("/users", postUser)
	router.Run("localhost:8080")
	router.POST("/login", loginHandler)
}

// insertar nuevos hoteles

func postHotels(context *gin.Context) {
	var newHotel Hotel

	err := context.BindJSON(&newHotel)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	createdHotel, err := CreateHotel(newHotel)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hotel"})
		return
	}

	context.JSON(http.StatusCreated, createdHotel)
}

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

func postReservations(context *gin.Context) {

	var newReservation Reservation
	if err := context.BindJSON(&newReservation); err != nil {
		return
	}

	//reservations = append(reservations, newReservation)
	context.IndentedJSON(http.StatusCreated, newReservation)
}

//confirma la reserva

func confirmReservation(reservation *Reservation) {

	reservation.IsConfirmed = true
}

func postUser(context *gin.Context) {
	var newUser User

	if err := context.BindJSON(&newUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	createdUser, err := CreateUser(newUser)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	context.JSON(http.StatusCreated, createdUser)
}

//-----------------------------------LOGIN-------------------------------------------

// loginHandler gestiona las solicitudes de inicio de sesión y autentica al usuario.

func loginHandler(context *gin.Context) {

	var user User
	// leer credenciales del formulario

	if err := context.ShouldBind(&user); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": "Credenciales inválidas"})
		return
	}

	// validar credenciales

	var userFound, err = Authenticate(user.EMail, user.Password)
	if userFound == nil {

		context.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	// inicio de sesión correcto, establecer cookie de sesión
	//--------------------------------------------------------------------------------------------------------------
	sessionToken, err := GenerateSessionToken()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token de sesión"})
		return
	}

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		HttpOnly: true,
		MaxAge:   3600,
		Path:     "/",
	}
	http.SetCookie(context.Writer, cookie)

	context.JSON(http.StatusOK, gin.H{"message": "Inicio de sesión exitoso"})

}

func getReservationByHotelId(context *gin.Context) {

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
