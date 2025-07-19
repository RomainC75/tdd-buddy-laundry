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

	store := db.GetConnection()
	reservations, err := (*store).ListReservations(suite.ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), len(reservations), 1)
	assert.NotNil(suite.T(), reservations[0])
	assert.Equal(suite.T(), email, reservations[0].Email)
	assert.Equal(suite.T(), pin, reservations[0].Pin)
}

func (suite *RepoTestSuite) TestCustomerRepositoryRead() {
	suite.T().Run("should return the data found", func(t *testing.T) {
		err := test.RunSQLFile(suite.pgContainer.Sql, filepath.Join("db", "migration", "tests", "init_for_read.sql"))
		assert.NoError(suite.T(), err)

		email := "bob1@email.com"
		reservationRepo := NewReservationRepo()
		foundReservation, err := reservationRepo.FindReservationByEmail(suite.ctx, email)

		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), "bob1@email.com", foundReservation.Email)
	})
	suite.T().Run("should return an error if no reservation has been found", func(t *testing.T) {
		reservationRepo := NewReservationRepo()
		email := "xxxxxx@email.com"
		_, err := reservationRepo.FindReservationByEmail(suite.ctx, "xxxxxx@email.com")

		assert.EqualError(suite.T(), err, fmt.Sprintf("no reservation found for email : %s", email))
	})
}

// !run the tests
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(RepoTestSuite))
}
