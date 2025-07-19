package repositories

import (
	"errors"
	"fmt"

	"laundry/core/business-logic/models"
)

type FakeReservationRepo struct {
	Reservations         []models.Reservation
	ShouldReturnError    bool
	ExpectedReservations []models.Reservation
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

func (frr *FakeReservationRepo) GetNextCloseReservations() ([]models.Reservation, error) {
	if frr.ShouldReturnError {
		return []models.Reservation{}, errors.New("error getting reservations")
	}
	return frr.ExpectedReservations, nil
}
