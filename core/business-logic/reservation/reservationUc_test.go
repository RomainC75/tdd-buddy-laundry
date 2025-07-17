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
		fakeReservationRepo := repositories.NewFakeReservationRepo()
		fakePinGenerator := providers.NewFakePinGenerator()
		fakePinGenerator.ExpectedPin = uCase.expectedPin
		fakeUuidGenerator := providers.NewFakeUuidGenerator()
		fakeUuidGenerator.ExpectedUuid = "1"

		reservationUC := NewReservationUC(fakeEmailProvider, fakePinGenerator, fakeReservationRepo, fakeUuidGenerator)
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
	fakeEmailProvider := providers.NewFakeEmailProvider()
	fakePinGenerator := providers.NewFakePinGenerator()
	fakeReservationRepo := repositories.NewFakeReservationRepo()
	fakeReservationRepo.ShouldReturnError = true
	fakeUuidGenerator := providers.NewFakeUuidGenerator()
	fakeUuidGenerator.ExpectedUuid = "112"

	reservationUC := NewReservationUC(fakeEmailProvider, fakePinGenerator, fakeReservationRepo, fakeUuidGenerator)
	reservationRequest := ReservationRequest{
		reservationDate: time.Now(),
		reservationTime: 45,
		email:           "johnDoe@email.com",
	}
	_, err := reservationUC.ReservationUseCase(reservationRequest)

	assert.EqualError(t, err, fmt.Sprintf("error trying to save the reservation : %s", "112"))
}
