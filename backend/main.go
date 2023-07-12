package main

import (
	. "backend/model"
	. "backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/hotels", getHotels)
	router.GET("/hotels/:id", getHotelById)
	router.POST("/hotels", postHotels)
	router.GET("/reservations", getReservations)
	router.GET("/reservations/:id", getReservationById)
	router.GET("/reservations/email/:email", getReservationByEmail)
	router.GET("/reservations/hotel/:id", getReservationByHotelId)
	router.POST("/reservations", postReservations)
	router.POST("/users", postUser)
	router.POST("/login", loginHandler)
	router.Run("localhost:8080")

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
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	id, err2 := strconv.ParseInt(context.Param("id"), 10, 64)
	if err2 != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

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
	id, err2 := strconv.ParseInt(context.Param("id"), 10, 64)
	if err2 != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	var reservation Reservation
	var err error

	reservation, err = GetReservationById(id)
	if err != nil {

		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, reservation)
}

/*
Frontend tiene que pasar un JSON como este:

	{
	    "checkin": "2023-01-01T15:04:05Z",
	    "checkout": "2023-01-04T15:04:05Z",
	    "idHotel": 1,
	    "email": "prueba@gmail.com"
	}
*/
func postReservations(context *gin.Context) {

	var newReservation Reservation

	if err := context.BindJSON(&newReservation); err != nil {

		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	var userExists, err = UserExists(newReservation.EMail)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo validar el usuario"})
		return
	}
	if !userExists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "El usuario no existe"})
		return
	}

	createdReservation, err := CreateReservation(newReservation)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver la nueva reserva en la respuesta

	context.IndentedJSON(http.StatusCreated, createdReservation)
}

func postUser(context *gin.Context) {

	var newUser User

	if err := context.BindJSON(&newUser); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var userExists, err = UserExists(newUser.EMail)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo validar el usuario"})
		return
	}
	if !userExists {
		createdUser, err := CreateUser(newUser)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario"})
			return
		}

		context.JSON(http.StatusCreated, createdUser)
	} else {
		context.JSON(http.StatusConflict, "El usuario ya existe")
	}

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
	id, err2 := strconv.ParseInt(context.Param("id"), 10, 64)
	if err2 != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	var reservations, err = GetReservationByHotelId(id)
	if err != nil {

		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, reservations)
}

func getReservationByEmail(context *gin.Context) {
	email := context.Param("email")

	var reservation, err = GetReservationByEmail(email)
	if err != nil {

		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, reservation)
}
