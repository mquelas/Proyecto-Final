package service

import (
	"fmt"
	"testing"

	//"context"
	"database/sql"
)

func TestUserExists_Exists(t *testing.T) {

	var userExists, err = UserExists("pepcito@gmail.com")

	if err != nil {

		t.Fatalf("Error al invocar UserExists: %v", err)
	}

	if !userExists {

		t.Fatalf("El usuario existe y no lo encontro")
	}
}

func TestUserExists_NotExists(t *testing.T) {

	var userExists, err = UserExists("luli_probando_test123@gmail.com")

	if err != nil {

		t.Fatalf("Error al invocar UserExists: %v", err)
	}

	if userExists {

		t.Fatalf("UserExits devuleve TRUE, pero el usuario no existe")
	}
}

func TestUserExists_Empty(t *testing.T) {

	var userExists, err = UserExists("")

	if err == nil {

		t.Fatalf("UserExists no falla con un email vacio: %v", err)
	}
	_ = userExists
}

func TestCreateUser_Success(t *testing.T) {

	user := User{
		EMail:    "peueba@example.com",
		Name:     "pepe",
		LastName: "gomez",
		Password: "123",
		Admin:    false,
	}

	createdUser, err := CreateUser(user)

	if err != nil {
		t.Fatalf("Error al invocar CreateUser: %v", err)
	}

	if createdUser.EMail != user.EMail ||
		createdUser.Name != user.Name ||
		createdUser.LastName != user.LastName ||
		createdUser.Password != user.Password ||
		createdUser.Admin != user.Admin {
		t.Fatalf("El usuario creado no coincide con el usuario de entrada")
	}
}

func TestCreateUser_Failure(t *testing.T) {

	user := User{
		EMail:    "test@example.com",
		Name:     "John",
		LastName: "Doe",
		Password: "password123",
		Admin:    false,
	}

	DataConnect = func() (*sql.DB, error) {
		return nil, fmt.Errorf("Error en la conexión a la base de datos")
	}
	_, err := CreateUser(user)

	if err == nil {
		t.Fatalf("CreateUser no falla cuando hay un error en la conexión a la base de datos")
	}
}
