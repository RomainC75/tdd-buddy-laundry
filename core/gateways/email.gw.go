package gateways

import (
	"laundry/core/business-logic/models"
)

type IEmail interface {
	send(reservationConfirmation models.Reservation) error
}
