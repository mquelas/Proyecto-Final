package model

// estructura reservas
type Reservation struct {
	ID          string `json:"id"`
	IdHotel     string `json:"idHotel"`
	CheckIn     string `json:"checkin"`
	CheckOut    string `json:"checkout"`
	IsConfirmed bool   `json:"isConfirmed"`
}
