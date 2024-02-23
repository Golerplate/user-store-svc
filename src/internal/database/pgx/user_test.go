package database_pgx

import (
	"context"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golerplate/pkg/constants"
	pkgerrors "github.com/golerplate/pkg/errors"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	entities_user_v1 "github.com/golerplate/user-store-svc/internal/entities/user/v1"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func Test_CreateUser(t *testing.T) {
	t.Run("ok - create user", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		mock.ExpectExec("INSERT INTO users").WithArgs(sqlmock.AnyArg(), "testuser", sqlmock.AnyArg(), "testuser@test.com", sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

		user, err := sqlxDB.CreateUser(context.Background(), &entities_user_v1.ServiceCreateUserRequest{
			ExternalID: "testuser",
			Email:      "testuser@test.com",
			Username:   "username",
		})
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.True(t, constants.User.IsValid(user.ID))
		assert.Equal(t, "testuser", user.ExternalID)
		assert.Equal(t, "username", user.Username)
		assert.Equal(t, "testuser@test.com", user.Email)
		assert.False(t, user.CreatedAt.IsZero())
		assert.False(t, user.CreatedAt.IsZero())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("nok - create user", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		mock.ExpectExec("INSERT INTO users").WithArgs(sqlmock.AnyArg(), "testuser", sqlmock.AnyArg(), "testuser@test.com", sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(pkgerrors.NewInternalServerError("error"))

		user, err := sqlxDB.CreateUser(context.Background(), &entities_user_v1.ServiceCreateUserRequest{
			ExternalID: "testuser",
			Email:      "testuser@test.com",
			Username:   "username",
		})
		assert.Nil(t, user)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func Test_GetUserByEmail(t *testing.T) {
	t.Run("ok - get user by email", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		userID := constants.GenerateDataPrefixWithULID(constants.User)

		rows := sqlmock.NewRows([]string{"id", "external_id", "username", "email", "created_at", "updated_at"}).
			AddRow(userID, "testuser", "username", "testuser@test.com", time.Now().String(), time.Now().String())

		mock.ExpectQuery("SELECT id, external_id, username, email, created_at, updated_at FROM users WHERE email = $1;").WithArgs("testuser@test.com").WillReturnError(nil).WillReturnRows(rows)

		user, err := sqlxDB.GetUserByEmail(context.Background(), "testuser@test.com")
		assert.NotNil(t, user)
		assert.NoError(t, err)

		// assert.True(t, constants.User.IsValid(user.ID))
		// assert.Equal(t, "testuser", user.ExternalID)
		// assert.Equal(t, "username", user.Username)
		// assert.Equal(t, "testuser@test.com", user.Email)
		// assert.False(t, user.CreatedAt.IsZero())
		// assert.False(t, user.UpdatedAt.IsZero())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func Test_UpdateUsername(t *testing.T) {
	t.Run("ok - update username", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		userID := constants.GenerateDataPrefixWithULID(constants.User)

		rows := sqlmock.NewRows([]string{"id", "external_id", "username", "email", "created_at", "updated_at"}).
			AddRow(userID, "testuser", "username", "testuser@test.com", time.Now(), time.Now())

		mock.ExpectQuery("UPDATE users SET username = $1, updated_at = $2 WHERE id = $3 RETURNING id, external_id, username, email, created_at, updated_at;").WithArgs("username", AnyTime{}, userID).WillReturnRows(rows)

		user, err := sqlxDB.UpdateUsername(context.Background(), userID, "username")
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.True(t, constants.User.IsValid(user.ID))
		assert.Equal(t, "testuser", user.ExternalID)
		assert.Equal(t, "username", user.Username)
		assert.Equal(t, "testuser@test.com", user.Email)
		assert.False(t, user.CreatedAt.IsZero())
		assert.False(t, user.UpdatedAt.IsZero())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	// t.Run("nok - update username", func(t *testing.T) {
	// 	db, mock, err := sqlmock.New()
	// 	assert.NoError(t, err)
	// 	defer db.Close()

	// 	sqlxDB := &dbClient{
	// 		connection: sqlx.NewDb(db, "sqlmock"),
	// 	}

	// 	userID := constants.GenerateDataPrefixWithULID(constants.User)

	// 	mock.ExpectQuery("UPDATE users SET username = $1, updated_at = $2 WHERE id = $3 RETURNING id, external_id, username, email, created_at, updated_at;").WithArgs("username", AnyTime{}, userID).WillReturnRows(sqlmock.NewRows([]string{"id", "external_id", "username", "email", "created_at", "updated_at"}))

	// 	user, err := sqlxDB.UpdateUsername(context.Background(), userID, "username")
	// 	assert.Nil(t, user)
	// 	assert.Error(t, err)

	// 	if err := mock.ExpectationsWereMet(); err != nil {
	// 		t.Errorf("there were unfulfilled expectations: %s", err)
	// 	}
	// })
}
