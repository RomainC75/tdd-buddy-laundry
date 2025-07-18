//go:build integration
// +build integration

package repositories

import (
	"context"
	"fmt"
	db "laundry/adapters/secondary/repositories/sqlc/db/sqlc"
	"laundry/config"
	"laundry/core/business-logic/models"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestCustomerRepository(t *testing.T) {
	ctx := context.Background()
	config.Set("postgres", "postgres", "localhost", 5432, "testdb")
	cfg := config.Get()

	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithInitScripts(filepath.Join("db", "migration", "init_test.sh")),
		postgres.WithDatabase(cfg.Db.Name),
		postgres.WithUsername(cfg.Db.Username),
		postgres.WithPassword(cfg.Db.Password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		fmt.Println("-----------> ", err.Error())
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)
	db.ConnectWithString(connStr)

	// customerRepo, err := NewRepository(ctx, connStr)
	// assert.NoError(t, err)
	reservationRepo := NewReservationRepo()

	newUuid := uuid.New()
	now := time.Now()
	var reservationTime int32 = 54
	email := "bob@email.com"
	pin := "abc1"
	machineNum := "a10"

	err = reservationRepo.CreateReservation(ctx, models.Reservation{
		Id:              newUuid,
		ReservationDate: now,
		ReservationTime: reservationTime,
		Email:           email,
		Pin:             pin,
		MachineNum:      machineNum,
	})

	assert.NoError(t, err)
	// assert.NotNil(t, c)

	store := db.GetConnection()
	reservations, err := (*store).ListReservations(ctx)

	assert.NoError(t, err)
	assert.Equal(t, len(reservations), 1)
	assert.NotNil(t, reservations[0])
	assert.Equal(t, email, reservations[0].Email)
	assert.Equal(t, pin, reservations[0].Pin)
}
