package repositories

import (
	"laundry/core/business-logic/models"
)

type FakeEmailRepo struct {
	Sent []string
}

func (fer *FakeEmailRepo) Send(reservationConfirmation models.Reservation) error {
	fer.Sent = []string{
		reservationConfirmation.ReservationId,
	}
	return nil
}
