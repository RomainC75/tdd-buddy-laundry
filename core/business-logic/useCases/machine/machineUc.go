package machine

import (
	"laundry/core/business-logic/models"
	"laundry/core/gateways"
)

type TNextReservedMachine struct {
	reservation models.Reservation
	isReserved  bool
}

type MachineUc struct {
	reservationRepo gateways.IReservationRepo
	machineProvider gateways.IMachineProvider
}

func NewMachineUc(reservationRepo gateways.IReservationRepo, machineProvider gateways.IMachineProvider) *MachineUc {
	return &MachineUc{
		reservationRepo: reservationRepo,
		machineProvider: machineProvider,
	}
}

func (muc *MachineUc) ReserveMachines() ([]TNextReservedMachine, error) {
	nextReservations, err := muc.reservationRepo.GetNextCloseReservations()
	if err != nil {
		return []TNextReservedMachine{}, err
	}

	nextReservedMachines := []TNextReservedMachine{}
	for _, reservation := range nextReservations {
		isReserved := muc.machineProvider.Lock(reservation.Id, reservation.ReservationDate, reservation.ReservationTime, reservation.MachineNum, reservation.Pin)
		nextReservedMachines = append(nextReservedMachines, TNextReservedMachine{
			reservation: reservation,
			isReserved:  isReserved,
		})
	}

	return nextReservedMachines, nil
}
