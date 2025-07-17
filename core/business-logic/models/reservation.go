package models

import "time"

type Reservation struct {
	ReservationDate time.Time
	ReservationTime int
	Email           string
	Pin             string
	ReservationId   string
}
