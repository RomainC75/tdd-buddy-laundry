package gateways

import "laundry/core/business-logic/models"

type IReservationRepo interface {
	Save(reservation models.Reservation) error
}
