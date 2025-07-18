package providers

import (
	"fmt"
	"laundry/core/business-logic/models"

	"github.com/google/uuid"
)

type FakeEmailProvider struct {
	Sent            []uuid.UUID
	ShouldSendError bool
}

func NewFakeEmailProvider() *FakeEmailProvider {
	return &FakeEmailProvider{}
}

func (fer *FakeEmailProvider) Send(reservationConfirmation models.Reservation) error {
	if fer.ShouldSendError {
		return fmt.Errorf("error trying to send email : %s", reservationConfirmation.Id)
	}
	fer.Sent = []uuid.UUID{
		reservationConfirmation.Id,
	}
	return nil
}
