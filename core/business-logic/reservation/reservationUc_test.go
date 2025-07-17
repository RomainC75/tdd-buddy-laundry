package reservation

import (
	"fmt"
	"laundry/adapters/secondary/providers"
	"laundry/adapters/secondary/repositories"
	"laundry/core/business-logic/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var uCases = []struct {
	reservationRequest    ReservationRequest
	expectedPin           string
	expectedReservationId string
}{
	{
		ReservationRequest{time.Now(), 45, "john.doe@email.com"}, "2345", "1",
	},
}

func TestReservation(t *testing.T) {
	for _, uCase := range uCases {
		fakeEmailProvider := providers.NewFakeEmailProvider()
		fakePinGenerator := providers.NewFakePinGenerator()
		fakePinGenerator.ExpectedPin = uCase.expectedPin
		fakeReservationRepo := repositories.NewFakeReservationRepo()

		reservationUC := NewReservationUC(fakeEmailProvider, fakePinGenerator, fakeReservationRepo)
		confirmation, err := reservationUC.ReservationUseCase(uCase.reservationRequest)

		assert.Nil(t, err)
		expectedReservation := models.Reservation{
			ReservationTime: uCase.reservationRequest.reservationTime,
			ReservationDate: uCase.reservationRequest.reservationDate,
			Email:           uCase.reservationRequest.email,
			Pin:             uCase.expectedPin,
			ReservationId:   "1",
		}
		assert.Equal(t, confirmation, expectedReservation)
		assert.Equal(t, len(fakeReservationRepo.Reservations), 1)
		assert.Equal(t, fakeReservationRepo.Reservations[0], expectedReservation)

		assert.Equal(t, fakeEmailProvider.Sent, []string{"1"})
	}
}

func TestRepoError(t *testing.T) {
	for _, uCase := range uCases {
		fakeEmailProvider := providers.NewFakeEmailProvider()
		fakePinGenerator := providers.NewFakePinGenerator()
		fakePinGenerator.ExpectedPin = uCase.expectedPin
		fakeReservationRepo := repositories.NewFakeReservationRepo()
		fakeReservationRepo.ShouldReturnError = true

		reservationUC := NewReservationUC(fakeEmailProvider, fakePinGenerator, fakeReservationRepo)
		_, err := reservationUC.ReservationUseCase(uCase.reservationRequest)

		assert.Error(t, err, fmt.Errorf("error trying to save the reservation : %s", uCase.expectedReservationId))
	}
}
