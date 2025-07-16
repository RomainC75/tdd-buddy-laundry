package reservation

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReservation(t *testing.T) {
	request := reservationRequest{
		reservationDate: time.Now(),
		reservationTime: 45,
		email:           "john.doe@email.com",
	}

	reservationUseCase := &reservationUC{}
	confirmation, err := reservationUseCase.ReservationUseCase(request)
	assert.Nil(t, err)

	assert.Equal(t, confirmation, reservationConfirmation{
		reservationRequest: request,
		pin:                "1234",
		reservationId:      "1",
	})
}
