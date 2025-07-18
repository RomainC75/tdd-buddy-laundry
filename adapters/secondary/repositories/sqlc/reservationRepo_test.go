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

	err = reservationRepo.CreateReservation(ctx, models.Reservation{
		Id:              uuid.New(),
		ReservationDate: time.Now(),
		ReservationTime: 45,
		Email:           "bob",
		Pin:             "abc1",
		MachineNum:      "a10",
	})
	if err != nil {
		fmt.Println("-----> create reservation ", err.Error())
	}
	assert.NoError(t, err)
	// assert.NotNil(t, c)

	store := db.GetConnection()
	reservations, err := (*store).ListReservations(ctx)
	if err != nil {
		fmt.Println("-----> reservation error : ", err.Error())
	}
	fmt.Println("--> reservations : ", reservations)

	// customer, err := customerRepo.GetCustomerByEmail(ctx, "henry@gmail.com")
	// assert.NoError(t, err)
	// assert.NotNil(t, customer)
	// assert.Equal(t, "Henry", customer.Name)
	// assert.Equal(t, "henry@gmail.com", customer.Email)
}
