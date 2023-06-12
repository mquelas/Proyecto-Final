package model

import (
	"time"
)

type Ocuppancy struct {
	ID            string    `json:"id"`
	Date          time.Time `json:"date"`
	IdReservation int64     `json:"idReservation"`
}
