package gateways

import (
	"time"

	"github.com/google/uuid"
)

type IMachineProvider interface {
	Lock(reservationId uuid.UUID, reservationDate time.Time, reservationTime int32, machineNum string, pin string) bool
	//   void Unlock(string reservationId);
}
