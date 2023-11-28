package repository_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/google/uuid"
	"github.com/ronnachate/foodstore/food-api/domain"
	"github.com/ronnachate/foodstore/food-api/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type OrderTestSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository domain.OrderRepository
}

func (s *OrderTestSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})

	s.DB, _ = gorm.Open(dialector, &gorm.Config{})
	s.repository = repository.NewOrderRepository(s.DB)
}

func (s *OrderTestSuite) AfterOrderTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestOrderInit(t *testing.T) {
	suite.Run(t, new(OrderTestSuite))
}

func (s *OrderTestSuite) Test_repository_NewOrder() {

	s.T().Run("error", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec("INSERT INTO \"orders\" (.+) VALUES (.+)").WillReturnError(sql.ErrNoRows)
		s.mock.ExpectRollback()
		var reqOrder domain.Order
		_, err := s.repository.NewOrder(context.Background(), &reqOrder)
		assert.Error(t, err)
	})
}

func (s *OrderTestSuite) Test_repository_GetByID() {
	var (
		id = uuid.UUID{}
	)
	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "orders" WHERE id = $1 ORDER BY "orders"."id" LIMIT 1`)).
			WithArgs(id.String()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).
				AddRow(id.String()))

		res, err := s.repository.GetByID(context.Background(), id.String())

		require.NoError(t, err)
		assert.Equal(t, id, res.ID)
	})

	s.T().Run("error", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "orders" WHERE id = $1 ORDER BY "orders"."id" LIMIT 1`)).
			WithArgs(id.String()).
			WillReturnError(sql.ErrNoRows)

		_, err := s.repository.GetByID(context.Background(), id.String())

		assert.Error(t, err)
	})
}
