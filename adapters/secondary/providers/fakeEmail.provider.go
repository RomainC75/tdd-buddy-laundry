package providers

import (
	"fmt"
	"laundry/core/business-logic/models"
)

type FakeEmailProvider struct {
	Sent            []string
	ShouldSendError bool
}

func NewFakeEmailProvider() *FakeEmailProvider {
	return &FakeEmailProvider{}
}

func (fer *FakeEmailProvider) Send(reservationConfirmation models.Reservation) error {
	if fer.ShouldSendError {
		return fmt.Errorf("error trying to send email : %s", reservationConfirmation.ReservationId)
	}
	fer.Sent = []string{
		reservationConfirmation.ReservationId,
	}
	return nil
}
