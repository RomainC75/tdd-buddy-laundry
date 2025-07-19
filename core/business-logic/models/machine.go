package models

import "github.com/google/uuid"

type Machine struct {
	MachineNum           string
	CurrentReservationId uuid.UUID
	NextReservationId    uuid.UUID
}
