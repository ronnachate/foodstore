package repository_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/ronnachate/foodstore/food-api/domain"
	"github.com/ronnachate/foodstore/food-api/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TestSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository domain.OrderDiscountRepository
}

func (s *TestSuite) SetupSuite() {
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
	s.repository = repository.NewOrderDiscountRepository(s.DB)
}

func (s *TestSuite) AfterOrderDiscountTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestOrderDiscountInit(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) Test_repository_GetByType() {
	var (
		typeID = uint64(99)
	)
	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "order_discounts" WHERE type = $1 ORDER BY "order_discounts"."id" LIMIT 1`)).
			WithArgs(typeID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "type"}).
				AddRow(1, typeID))

		res, err := s.repository.GetByType(context.Background(), typeID)

		require.NoError(t, err)
		assert.Equal(t, typeID, res.Type)
	})

	s.T().Run("error", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "order_discounts" WHERE type = $1 ORDER BY "order_discounts"."id" LIMIT 1`)).
			WithArgs(typeID).
			WillReturnError(sql.ErrNoRows)

		_, err := s.repository.GetByType(context.Background(), typeID)

		assert.Error(t, err)
	})
}
