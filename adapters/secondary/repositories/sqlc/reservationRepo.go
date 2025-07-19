package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	db "laundry/adapters/secondary/repositories/sqlc/db/sqlc"
	"laundry/core/business-logic/models"
	// db "laundry/adapter/secondary/repositories/db"
)

type ReservationRepository struct {
	Store *db.Store
}

func NewReservationRepo() *ReservationRepository {
	return &ReservationRepository{
		Store: db.GetConnection(),
	}
}

func (reservationRepo *ReservationRepository) CreateReservation(ctx context.Context, arg models.Reservation) error {
	now := time.Now()
	newReservation := db.CreateReservationParams{
		ID:              arg.Id,
		ReservationDate: arg.ReservationDate,
		ReservationTime: arg.ReservationTime,
		Email:           arg.Email,
		Pin:             arg.Pin,
		MachineNum:      arg.MachineNum,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	_, err := (*reservationRepo.Store).CreateReservation(ctx, newReservation)
	return err
}

func (reservationRepo *ReservationRepository) FindReservationByEmail(ctx context.Context, email string) (models.Reservation, error) {
	foundResa, err := (*reservationRepo.Store).GetReservationByEmail(ctx, email)
	if err != nil && err == sql.ErrNoRows {
		return models.Reservation{}, fmt.Errorf("no reservation found for email : %s", email)
	}
	if err != nil {
		return models.Reservation{}, err
	}

	return models.ReservationFromSnapshot(foundResa.ID, foundResa.ReservationDate, foundResa.ReservationTime, foundResa.Email, foundResa.Pin, foundResa.MachineNum), nil
}
