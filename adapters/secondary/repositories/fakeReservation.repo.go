package repositories

import (
	"fmt"
	"laundry/core/business-logic/models"
)

type FakeReservationRepo struct {
	Reservations      []models.Reservation
	ShouldReturnError bool
}

func NewFakeReservationRepo() *FakeReservationRepo {
	return &FakeReservationRepo{}
}

func (frr *FakeReservationRepo) Save(reservation models.Reservation) error {
	if frr.ShouldReturnError {
		return fmt.Errorf("error trying to save the reservation : %s", reservation.Id)
	}
	frr.Reservations = append(frr.Reservations, reservation)
	return nil
}
