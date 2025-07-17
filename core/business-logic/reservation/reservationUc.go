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
	emailRepo    gateways.IEmail
	pinGenerator gateways.IPinGenerator
}

func NewReservationUC(emailProvider gateways.IEmail, pinGenerator gateways.IPinGenerator) *ReservationUC {
	return &ReservationUC{
		emailRepo:    emailProvider,
		pinGenerator: pinGenerator,
	}
}

func (rus *ReservationUC) ReservationUseCase(reservationRequest ReservationRequest) (models.Reservation, error) {
	// need to be calculated
	newReservationId := "1"
	reservation := models.Reservation{
		ReservationTime: reservationRequest.reservationTime,
		ReservationDate: reservationRequest.reservationDate,
		Email:           reservationRequest.email,
		Pin:             rus.pinGenerator.Generate(),
		ReservationId:   newReservationId,
	}
	rus.emailRepo.Send(reservation)

	return reservation, nil
}
