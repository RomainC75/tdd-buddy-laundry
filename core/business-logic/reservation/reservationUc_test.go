package reservation

import (
	"laundry/adapters/secondary/providers"
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

	fakeEmailProvider := providers.NewFakeEmailProvider()
	fakePinGenerator := providers.NewFakePinGenerator()
	fakePinGenerator.ExpectedPin = "9876"

	reservationUC := NewReservationUC(fakeEmailProvider, fakePinGenerator)
	confirmation, err := reservationUC.ReservationUseCase(request)

	assert.Nil(t, err)

	assert.Equal(t, confirmation, models.Reservation{
		ReservationTime: request.reservationTime,
		ReservationDate: request.reservationDate,
		Email:           request.email,
		Pin:             "9876",
		ReservationId:   "1",
	})

	assert.Equal(t, fakeEmailProvider.Sent, []string{"1"})
}
