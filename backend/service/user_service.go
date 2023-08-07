package service

import (
	. "backend/data_access"
	. "backend/model"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func UserExists(email string) (bool, error) {
	if email == "" {
		return false, fmt.Errorf("El email no puede ser vacio")
	}

	var err error
	var db *sql.DB

	db, err = DataConnect()

	if err != nil {

		return false, fmt.Errorf("Error al conectar a la base de datos: %v", err)
	}

	defer db.Close()

	row := db.QueryRow("SELECT count(1) FROM user WHERE email = ? ", email)

	var foundUsers int64
	if err = row.Scan(&foundUsers); err != nil {

		if err == sql.ErrNoRows {

			return false, fmt.Errorf("Usuario no encontrado")
		}
		return false, fmt.Errorf("Error al obtener el usuario de la base de datos: %v", err)
	}
	if foundUsers == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func CreateUser(user User) (User, error) {

	var err error
	var db *sql.DB

	db, err = DataConnect()

	if err != nil {

		return user, fmt.Errorf("createUser %q", err)
	}

	defer db.Close()

	insertResult, err := db.ExecContext(

		context.Background(),
		"INSERT INTO user (email, name, lastname, password, admin) VALUES (?, ?, ?, ?, ?)",
		user.EMail, user.Name, user.LastName, user.Password, user.Admin,
	)
	_ = insertResult

	if err != nil {
		return user, fmt.Errorf("createUser %q", err)
	}

	return user, nil
}

func Authenticate(email string, password string) (*User, error) {

	var user *User
	var err error
	var db *sql.DB

	db, err = DataConnect()

	if err != nil {

		return nil, fmt.Errorf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT password FROM user WHERE email = ?", email)

	if err != nil {

		return nil, fmt.Errorf("Error al obtener el usuario de la base de datos: %v", err)
	}
	defer rows.Close()

	if rows.Next() {

		err = rows.Scan(&user)

		if err != nil {

			return nil, fmt.Errorf("Error al escanear la contraseña del usuario: %v", err)

		}
	} else {

		return nil, fmt.Errorf("Usuario no encontrado")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {

		return nil, fmt.Errorf("Credenciales incorrectas")
	}

	return user, nil
}

// Generar un token de sesión

func GenerateSessionToken() (string, error) {

	// Definir la clave secreta para firmar el token

	secretKey := "mi_clave_secreta"

	// Crear una estructura de reclamaciones (claims)
	claims := jwt.MapClaims{

		"exp": time.Now().Add(time.Hour * 1).Unix(), // Tiempo de expiración del token
	}

	// Crear el token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token con la clave secreta
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
