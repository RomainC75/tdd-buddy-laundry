package repositories

import (
	"context"
	db "laundry/adapters/secondary/repositories/sqlc/db/sqlc"
	"laundry/core/business-logic/models"

	"time"
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

// func (userRepo *ReservationRepository) FindUserByEmail(ctx context.Context, email string) (db.User, error) {
// 	foundUser, err := (*userRepo.Store).GetUserByEmail(ctx, email)
// 	if err != nil {
// 		return db.User{}, err
// 	}
// 	return foundUser, nil
// }
