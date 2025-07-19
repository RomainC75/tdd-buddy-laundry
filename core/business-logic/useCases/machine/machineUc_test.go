package machine

import (
	"testing"
	"time"

	"laundry/adapters/secondary/providers"
	"laundry/adapters/secondary/repositories"
	"laundry/core/business-logic/models"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MachineUcTestSuite struct {
	suite.Suite
}

func (suite *MachineUcTestSuite) TestMachine() {
	suite.T().Run("should lock a machine", func(t *testing.T) {
		reservations := []models.Reservation{
			{
				Id:              uuid.New(),
				ReservationDate: time.Now(),
				ReservationTime: 43,
				Email:           "bob@email.com",
				Pin:             "a234",
				MachineNum:      "a10",
			},
		}

		fakeReservationRepo := repositories.NewFakeReservationRepo()
		fakeReservationRepo.ExpectedReservations = reservations

		fakeMachineProvider := providers.NewFakeMachineProvider()
		fakeMachineProvider.ExpectedResponse = true

		fakeMachineUc := NewMachineUc(fakeReservationRepo, fakeMachineProvider)
		nextReservations, err := fakeMachineUc.ReserveMachines()
		assert.Nil(t, err)
		assert.Equal(t, nextReservations, []TNextReservedMachine{
			{reservation: reservations[0], isReserved: true},
		})
	})
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(MachineUcTestSuite))
}
