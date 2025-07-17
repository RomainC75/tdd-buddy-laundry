package models

import "time"

type Reservation struct {
	Id              string
	ReservationDate time.Time
	ReservationTime int
	Email           string
	Pin             string
}
