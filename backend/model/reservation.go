package model

import (
	"time"
)

// estructura reservas
type Reservation struct {
	ID          int64     `json:"id"`
	CheckIn     time.Time `json:"checkin"`
	CheckOut    time.Time `json:"checkout"`
	IdHotel     int64     `json:"idHotel"`
	EMail       string    `json:"email"`
	IsConfirmed bool      `json:"isConfirmed"`
}
