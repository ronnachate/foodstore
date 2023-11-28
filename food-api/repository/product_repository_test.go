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

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository domain.ProductRepository
}

func (s *Suite) SetupSuite() {
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
	s.repository = repository.NewProductRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_repository_GetUsers() {
	var (
		id   = uuid.UUID{}
		name = "test-product"
	)
	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "products" WHERE id IN ($1)`)).
			WithArgs(id.String()).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
				AddRow(id.String(), name))

		users, err := s.repository.GetProducts(context.Background(), []uuid.UUID{id})

		require.NoError(t, err)
		assert.Equal(t, 1, len(users))
		assert.Equal(t, id, users[0].ID)
		assert.Equal(t, name, users[0].Name)
	})

	s.T().Run("error", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "products" WHERE id IN ($1)`)).
			WithArgs(id.String()).
			WillReturnError(sql.ErrNoRows)

		_, err := s.repository.GetProducts(context.Background(), []uuid.UUID{id})

		assert.Error(t, err)
	})
}
