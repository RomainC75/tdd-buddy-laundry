package reservation

import "time"

type reservationRequest struct {
	reservationDate time.Time
	reservationTime int
	email           string
}

type reservationConfirmation struct {
	reservationRequest
	pin           string
	reservationId string
}

type reservationUC struct {
}

func (rus *reservationUC) ReservationUseCase(reservationRequest reservationRequest) (reservationConfirmation, error) {
	return reservationConfirmation{
		reservationRequest: reservationRequest,
		pin:                "1234",
		reservationId:      "1",
	}, nil
}
