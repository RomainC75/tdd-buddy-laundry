package providers

import (
	"time"

	"laundry/core/business-logic/models"

	"github.com/google/uuid"
)

var Machines = []models.Machine{
	{"a10", uuid.UUID{}, uuid.UUID{}},
}

type FakeMachineProvider struct {
	ExpectedResponse bool
}

func NewFakeMachineProvider() *FakeMachineProvider {
	return &FakeMachineProvider{}
}

func (fm *FakeMachineProvider) Lock(reservationId uuid.UUID, reservationDate time.Time, reservationTime int32, machineNum string, pin string) bool {
	return fm.ExpectedResponse
}
