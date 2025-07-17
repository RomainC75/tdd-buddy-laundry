package providers

import (
	"laundry/core/business-logic/models"
)

type FakeEmailProvider struct {
	Sent []string
}

func NewFakeEmailProvider() *FakeEmailProvider {
	return &FakeEmailProvider{}
}

func (fer *FakeEmailProvider) Send(reservationConfirmation models.Reservation) error {
	fer.Sent = []string{
		reservationConfirmation.ReservationId,
	}
	return nil
}
