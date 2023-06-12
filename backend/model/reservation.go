package model

import (
	"time"
)

// estructura reservas
type Reservation struct {
	ID       int64     `json:"id"`
	CheckIn  time.Time `json:"checkin" time_format:"2006-01-02"`
	CheckOut time.Time `json:"checkout" time_format:"2006-01-02"`
	IdHotel  int64     `json:"idHotel"`
	EMail    string    `json:"email"`
}
