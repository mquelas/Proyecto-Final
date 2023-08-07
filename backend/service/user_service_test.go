package service

import (
	. "backend/data_access"
	"context"
	"database/sql"
	"testing"

	. "backend/model"
)

var emailNewUser = "prueba@example.com"
var emailExistingUser = "pepcito@gmail.com"

func TestUserExists_Exists(t *testing.T) {
	//Pre-test ==================
	deleteTestData(t, emailExistingUser)
	createTestData__TestUserExists_Exists(t)

	var userExists, err = UserExists(emailExistingUser)

	if err != nil {
		t.Fatalf("Error al invocar UserExists: %v", err)
	}

	if !userExists {
		t.Fatalf("El usuario existe y no lo encontro")
	}
	//Post-test ==================
	deleteTestData(t, emailExistingUser)
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

	if userExists {

		t.Fatalf("UserExists no falla con un email vacio: %v", err)
	}
	_ = userExists
}

func TestCreateUser_Success(t *testing.T) {
	//Pre-test ====================================
	deleteTestData(t, emailNewUser)
	//Fin pre-test =====================================

	var user = User{
		EMail:    emailNewUser,
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

	//Post-test ==================
	deleteTestData(t, emailNewUser)
	//Fin post-test
}

func TestCreateUser_FailureNoEmail(t *testing.T) {
	user := User{
		EMail:    "",
		Name:     "John",
		LastName: "Doe",
		Password: "password123",
		Admin:    false,
	}

	_, err := CreateUser(user)

	if err == nil {
		t.Fatalf("CreateUser no falla cuando hay un error en la conexión a la base de datos")
	}
}

func TestCreateUser_FailureUserExists(t *testing.T) {
	user := User{
		EMail:    emailExistingUser,
		Name:     "John",
		LastName: "Doe",
		Password: "password123",
		Admin:    false,
	}

	_, err := CreateUser(user)

	if err == nil {
		t.Fatalf("CreateUser no falla cuando se quiere crear un usuario que ya existe")
	}
}

func deleteTestData(t *testing.T, email string) {
	var db *sql.DB
	var err error

	db, err = DataConnect()

	if err != nil {
		t.Fatalf("Error al Conectarse a la DB %v", err)
	}
	defer db.Close()

	queryResult, err := db.ExecContext(
		context.Background(),
		"DELETE FROM user WHERE email = ?", email,
	)
	_ = queryResult

}

func createTestData__TestUserExists_Exists(t *testing.T) {
	user := User{
		EMail:    emailExistingUser,
		Name:     "John",
		LastName: "Doe",
		Password: "password123",
		Admin:    false,
	}

	_, err := CreateUser(user)
	if err != nil {
		t.Fatalf("Error al invocar CreateUser: %v", err)
	}
}
