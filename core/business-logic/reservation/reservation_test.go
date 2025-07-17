package reservation

import (
	"laundry/core/business-logic/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReservation(t *testing.T) {
	request := ReservationRequest{
		reservationDate: time.Now(),
		reservationTime: 45,
		email:           "john.doe@email.com",
	}

	reservationUseCase := &ReservationUC{}
	confirmation, err := reservationUseCase.ReservationUseCase(request)
	assert.Nil(t, err)

	assert.Equal(t, confirmation, models.Reservation{
		ReservationTime: request.reservationTime,
		ReservationDate: request.reservationDate,
		Email:           request.email,
		Pin:             "1234",
		ReservationId:   "1",
	})
}
