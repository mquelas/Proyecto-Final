package service

/*. "backend/data_access"
. "backend/model"
"database/sql"
"fmt" */

//var users []User

/*
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
}
*/
