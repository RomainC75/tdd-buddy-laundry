package gateways

import (
	"laundry/core/business-logic/models"
)

type IEmail interface {
	Send(reservationConfirmation models.Reservation) error
}
