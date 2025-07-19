package reservation

import (
	"laundry/core/business-logic/models"
	"laundry/core/gateways"
	"laundry/helper"
	"time"
)

var machines = []string{
	"a1",
	"a2",
	"b1",
	"bA",
}

type ReservationRequest struct {
	reservationDate time.Time
	reservationTime int32
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
		MachineNum:      helper.GetRandomInArray[string](machines),
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
