package models

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	Id              uuid.UUID
	ReservationDate time.Time
	ReservationTime int32
	Email           string
	Pin             string
	MachineNum      string
}

func ReservationFromSnapshot(id uuid.UUID, resDate time.Time, resTime int32, email string, pin string, machineN string) Reservation {
	return Reservation{
		Id:              id,
		ReservationDate: resDate,
		ReservationTime: resTime,
		Email:           email,
		Pin:             pin,
		MachineNum:      machineN,
	}
}
