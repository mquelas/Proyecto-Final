package model

// estructura Usuario
type User struct {
	EMail    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}
