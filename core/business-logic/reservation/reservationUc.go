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
	emailRepo       gateways.IEmail
	pinGenerator    gateways.IPinGenerator
	reservationRepo gateways.IReservationRepo
}

func NewReservationUC(emailProvider gateways.IEmail, pinGenerator gateways.IPinGenerator, reservationRepo gateways.IReservationRepo) *ReservationUC {
	return &ReservationUC{
		emailRepo:       emailProvider,
		pinGenerator:    pinGenerator,
		reservationRepo: reservationRepo,
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
	err := rus.reservationRepo.Save(reservation)
	if err != nil {
		return models.Reservation{}, err
	}
	err = rus.emailRepo.Send(reservation)
	if err != nil {
		return models.Reservation{}, err
	}
	return reservation, nil
}
