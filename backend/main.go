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
func postHotels(context *gin.Context) {

	var newHotel Hotel
	if err := context.BindJSON(&newHotel); err != nil {
		return
	}

	//hotels = append(hotels, newHotel)
	context.IndentedJSON(http.StatusCreated, newHotel)
}

func getHotels(context *gin.Context) {
	var hotels []Hotel
	var err error

	hotels, err = GetHotels()
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}
	context.IndentedJSON(http.StatusOK, hotels)
}

func getHotelById(context *gin.Context) {

	var hotels Hotel
	var err error
	id := context.Param("id")

	hotels, err = GetHotelById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, hotels)
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

/*var users []User

func main() {
    r := gin.Default()

    // Initialize session middleware
    store := cookie.NewStore([]byte("secret"))
    r.Use(sessions.Sessions("mysession", store))

    // Create user endpoint
    r.POST("/users", func(c *gin.Context) {
        var user User
        if err := c.BindJSON(&user); err != nil {
            c.AbortWithStatus(http.StatusBadRequest)
            return
        }

        // Check if user already exists
        for _, u := range users {
            if u.Username == user.Username {
                c.AbortWithStatus(http.StatusConflict)
                return
            }
        }

        users = append(users, user)
        c.Status(http.StatusCreated)
    })

    // Login endpoint
    r.POST("/login", func(c *gin.Context) {
        var user User
        if err := c.BindJSON(&user); err != nil {
            c.AbortWithStatus(http.StatusBadRequest)
            return
        }

        // Find user in list of registered users
        for _, u := range users {
            if u.Username == user.Username && u.Password == user.Password {

                // Set the user as authenticated in the session
                session := sessions.Default(c)
                session.Set("authenticated", true)
                session.Save()

                c.Status(http.StatusOK)
                return
            }
        }

        c.AbortWithStatus(http.StatusUnauthorized)
    })

    // Protected endpoint - requires authentication
    r.GET("/protected", AuthMiddleware(), func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "You are authenticated!"})
    })

    r.Run(":8080")
}

// Middleware to require authentication
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        session := sessions.Default(c)
        if auth, ok := session.Get("authenticated").(bool); !ok || !auth {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        // Continue processing request
        c.Next()
    }
}*/
