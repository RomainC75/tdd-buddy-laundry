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
	uuidGenerator   gateways.IUuidGenerator
}

func NewReservationUC(emailProvider gateways.IEmail, pinGenerator gateways.IPinGenerator, reservationRepo gateways.IReservationRepo, uuidGenerator gateways.IUuidGenerator) *ReservationUC {
	return &ReservationUC{
		emailRepo:       emailProvider,
		pinGenerator:    pinGenerator,
		reservationRepo: reservationRepo,
		uuidGenerator:   uuidGenerator,
	}
}

func (rus *ReservationUC) ReservationUseCase(reservationRequest ReservationRequest) (models.Reservation, error) {
	reservation := models.Reservation{
		ReservationTime: reservationRequest.reservationTime,
		ReservationDate: reservationRequest.reservationDate,
		Email:           reservationRequest.email,
		Pin:             rus.pinGenerator.Generate(),
		Id:              rus.uuidGenerator.Generate(),
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
