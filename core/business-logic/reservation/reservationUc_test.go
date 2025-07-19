package reservation

import (
	"fmt"
	"testing"
	"time"

	"laundry/adapters/secondary/providers"
	"laundry/adapters/secondary/repositories"
	"laundry/core/business-logic/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ReservationUcTestSuite struct {
	suite.Suite
}

var uCases = []struct {
	reservationRequest ReservationRequest
	expectedPin        string
	expectedId         string
}{
	{
		ReservationRequest{time.Now(), 45, "john.doe@email.com"}, "2345", "1",
	},
}

func (suite *ReservationUcTestSuite) TestReservation() {
	suite.T().Run("should return the data found", func(t *testing.T) {
		for _, uCase := range uCases {
			fakeEmailProvider := providers.NewFakeEmailProvider()
			fakeReservationRepo := repositories.NewFakeReservationRepo()
			fakePinGenerator := providers.NewFakePinGenerator()
			fakePinGenerator.ExpectedPin = uCase.expectedPin
			fakeUuidGenerator := providers.NewFakeUuidGenerator()
			uuidVar := uuid.MustParse("dda708b8-b760-45ff-89f5-cadcf14cc656")
			fakeUuidGenerator.ExpectedUuid = uuidVar

			reservationUC := NewReservationUC(fakeEmailProvider, fakePinGenerator, fakeReservationRepo, fakeUuidGenerator)
			confirmation, err := reservationUC.ReservationUseCase(uCase.reservationRequest)

			assert.Nil(t, err)
			expectedReservation := models.Reservation{
				ReservationTime: uCase.reservationRequest.reservationTime,
				ReservationDate: uCase.reservationRequest.reservationDate,
				Email:           uCase.reservationRequest.email,
				Pin:             uCase.expectedPin,
				Id:              uuidVar,
				// random parameter
				MachineNum: confirmation.MachineNum,
			}
			assert.Equal(t, confirmation, expectedReservation)
			assert.Equal(t, len(fakeReservationRepo.Reservations), 1)
			assert.Equal(t, fakeReservationRepo.Reservations[0], expectedReservation)

			assert.Equal(t, fakeEmailProvider.Sent, []uuid.UUID{uuidVar})
		}
	})

	suite.T().Run("should return a repo error", func(t *testing.T) {
		fakeEmailProvider := providers.NewFakeEmailProvider()
		fakePinGenerator := providers.NewFakePinGenerator()
		fakeReservationRepo := repositories.NewFakeReservationRepo()
		fakeReservationRepo.ShouldReturnError = true
		fakeUuidGenerator := providers.NewFakeUuidGenerator()
		uuidVar := uuid.MustParse("dda708b8-b760-45ff-89f5-cadcf14cc656")
		fakeUuidGenerator.ExpectedUuid = uuidVar

		reservationUC := NewReservationUC(fakeEmailProvider, fakePinGenerator, fakeReservationRepo, fakeUuidGenerator)
		reservationRequest := ReservationRequest{
			reservationDate: time.Now(),
			reservationTime: 45,
			email:           "johnDoe@email.com",
		}
		_, err := reservationUC.ReservationUseCase(reservationRequest)

		assert.EqualError(t, err, fmt.Sprintf("error trying to save the reservation : %s", uuidVar))
	})

	suite.T().Run("should return an email error", func(t *testing.T) {
		fakeEmailProvider := providers.NewFakeEmailProvider()
		fakeEmailProvider.ShouldSendError = true
		fakePinGenerator := providers.NewFakePinGenerator()
		fakeReservationRepo := repositories.NewFakeReservationRepo()
		fakeUuidGenerator := providers.NewFakeUuidGenerator()
		uuidVar := uuid.MustParse("dda708b8-b760-45ff-89f5-cadcf14cc656")
		fakeUuidGenerator.ExpectedUuid = uuidVar

		reservationUC := NewReservationUC(fakeEmailProvider, fakePinGenerator, fakeReservationRepo, fakeUuidGenerator)
		reservationRequest := ReservationRequest{
			reservationDate: time.Now(),
			reservationTime: 45,
			email:           "johnDoe@email.com",
		}
		_, err := reservationUC.ReservationUseCase(reservationRequest)

		assert.EqualError(t, err, fmt.Sprintf("error trying to send email : %s", uuidVar))
	})
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ReservationUcTestSuite))
}
