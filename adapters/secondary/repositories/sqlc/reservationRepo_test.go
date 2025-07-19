package repositories

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"testing"
	"time"

	db "laundry/adapters/secondary/repositories/sqlc/db/sqlc"
	"laundry/core/business-logic/models"
	"laundry/test"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepoTestSuite struct {
	suite.Suite
	pgContainer *test.PostgresContainer
	repo        *ReservationRepository
	ctx         context.Context
}

func (suite *RepoTestSuite) SetupSuite() {
	fmt.Println("=========================================")
	suite.ctx = context.Background()

	pgContainer, err := test.CreatePostgresContainer(suite.ctx)
	if err != nil {
		fmt.Println("-----------> ", err.Error())
		log.Fatal(err)
	}
	suite.pgContainer = pgContainer

	db.ConnectWithString(suite.pgContainer.ConnectionString)
}

func (suite *RepoTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *RepoTestSuite) TestCustomerRepository() {
	// t.Cleanup(func() {
	// 	if err := suite.pgContainer.Terminate(ctx); err != nil {
	// 		t.Fatalf("failed to terminate pgContainer: %s", err)
	// 	}
	// })

	// customerRepo, err := NewRepository(ctx, connStr)
	// assert.NoError(t, err)
	reservationRepo := NewReservationRepo()

	newUuid := uuid.New()
	now := time.Now()
	var reservationTime int32 = 54
	email := "bob@email.com"
	pin := "abc1"
	machineNum := "a10"

	err := reservationRepo.CreateReservation(suite.ctx, models.Reservation{
		Id:              newUuid,
		ReservationDate: now,
		ReservationTime: reservationTime,
		Email:           email,
		Pin:             pin,
		MachineNum:      machineNum,
	})

	assert.NoError(suite.T(), err)
	// assert.NotNil(suite.T(), c)

	store := db.GetConnection()
	reservations, err := (*store).ListReservations(suite.ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), len(reservations), 1)
	assert.NotNil(suite.T(), reservations[0])
	assert.Equal(suite.T(), email, reservations[0].Email)
	assert.Equal(suite.T(), pin, reservations[0].Pin)
}

func (suite *RepoTestSuite) TestCustomerRepositoryRead() {
	err := test.RunSQLFile(suite.pgContainer.Sql, filepath.Join("db", "migration", "tests", "init_for_read.sql"))
	if err != nil {
		fmt.Println("------->", err.Error())
	}
	assert.NoError(suite.T(), err)
	reservationRepo := NewReservationRepo()
	foundReservation, err := reservationRepo.FindReservationByEmail(suite.ctx, "bob1@email.com")

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), foundReservation.Email, "bob1@email.com")
}

// !run the tests
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(RepoTestSuite))
}
