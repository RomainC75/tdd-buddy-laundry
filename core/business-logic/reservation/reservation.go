package reservation

import (
	"laundry/core/business-logic/models"
	"laundry/core/gateways"
	"time"
)

type ReservationRequest struct {
	reservationDate time.Time
	reservationTime int
	email           string
}

type ReservationUC struct {
	emailRepo *gateways.IEmail
}

func NewReservationUC(emailRepo *gateways.IEmail) *ReservationUC {
	return &ReservationUC{
		emailRepo: emailRepo,
	}
}

func (rus *ReservationUC) ReservationUseCase(reservationRequest ReservationRequest) (models.Reservation, error) {
	// need to be calculated
	newPin := "1234"
	newReservationId := "1"

	return models.Reservation{
		ReservationTime: reservationRequest.reservationTime,
		ReservationDate: reservationRequest.reservationDate,
		Email:           reservationRequest.email,
		Pin:             newPin,
		ReservationId:   newReservationId,
	}, nil
}
