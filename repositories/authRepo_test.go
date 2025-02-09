package repositories

import (
	"database/sql"
	"errors"
	"github.com/AnkitDhawale/TodoListApp/domains"
	"github.com/AnkitDhawale/TodoListApp/dto"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

const (
	findByUserQuery = `SELECT * FROM users WHERE email = $1 AND password_hash = $2`
	findByEmail     = `SELECT * FROM users WHERE email = $1`
)

var (
	db             *sql.DB
	mock           sqlmock.Sqlmock
	err            error
	mockAuthRepoDb AuthRepoDb
	testUser       *dto.User
	testResp       *domains.User
)

func setup(t *testing.T) {
	db, mock, err = sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock connection: %v", err)
	}

	// Initialize AuthRepoDb
	mockAuthRepoDb = NewAuthRepoDb(db)

	testUser = &dto.User{
		Email:    "test@test.com",
		Password: "testPassword",
	}

	testResp = &domains.User{
		Id:           "111",
		Email:        "test@test.com",
		PasswordHash: "testPassword",
		CreatedAt:    time.Now(),
	}

	t.Cleanup(func() {
		_ = db.Close() // Ensuring db.Close() is called when the test ends
	})
}

func TestAuthRepoDb_FindUserBy(t *testing.T) {
	setup(t)

	tests := []struct {
		name         string
		inputUser    *dto.User
		mockQuery    func()
		expectedResp *domains.User
		expectedErr  error
	}{
		{
			name:      "success",
			inputUser: testUser,
			mockQuery: func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"email",
					"password_hash",
					"created_at",
				}).AddRow(
					testResp.Id,
					testResp.Email,
					testResp.PasswordHash,
					testResp.CreatedAt,
				)
				mock.ExpectQuery(regexp.QuoteMeta(findByUserQuery)).WithArgs(testUser.Email, testUser.Password).
					WillReturnRows(rows)
			},
			expectedResp: testResp,
			expectedErr:  nil,
		},
		{
			name:      "failure: Invalid credentials (no rows)",
			inputUser: testUser,
			mockQuery: func() {
				mock.ExpectQuery(regexp.QuoteMeta(findByUserQuery)).
					WithArgs(testUser.Email, testUser.Password).
					WillReturnError(sql.ErrNoRows)
			},
			expectedResp: nil,
			expectedErr:  ErrInvalidCredentials,
		},
		{
			name:      "failure: unexpected database error",
			inputUser: testUser,
			mockQuery: func() {
				mock.ExpectQuery(regexp.QuoteMeta(findByUserQuery)).
					WithArgs(testUser.Email, testUser.Password).
					WillReturnError(errors.New("db connection error"))
			},
			expectedResp: nil,
			expectedErr:  ErrUnexpected,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock behavior
			tt.mockQuery()

			// Execute the function
			res, err := mockAuthRepoDb.FindUserBy(tt.inputUser)

			// Assertions
			if tt.expectedErr != nil {
				assert.NotNil(t, err)
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, res, tt.expectedResp)
			}

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestAuthRepoDb_FindByEmail(t *testing.T) {
	setup(t)

	tests := []struct {
		name           string
		inputEmail     string
		mockQuery      func()
		expectedResult *domains.User
		expectedError  error
	}{
		{
			name:       "failure: invalid email",
			inputEmail: testUser.Email,
			mockQuery: func() {
				mock.ExpectQuery(regexp.QuoteMeta(findByEmail)).
					WithArgs(testUser.Email).
					WillReturnError(sql.ErrNoRows)
			},
			expectedResult: nil,
			expectedError:  ErrInvalidEmail,
		},
		{
			name:       "failure: unexpected database error",
			inputEmail: testUser.Email,
			mockQuery: func() {
				mock.ExpectQuery(regexp.QuoteMeta(findByEmail)).
					WithArgs(testUser.Email).
					WillReturnError(ErrUnexpected)
			},
			expectedResult: nil,
			expectedError:  ErrUnexpected,
		},
		{
			name:       "success",
			inputEmail: testUser.Email,
			mockQuery: func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"email",
					"password_hash",
					"created_at",
				}).AddRow(
					testResp.Id,
					testResp.Email,
					testResp.PasswordHash,
					testResp.CreatedAt,
				)

				mock.ExpectQuery(regexp.QuoteMeta(findByEmail)).
					WithArgs(testUser.Email).
					WillReturnRows(rows)
			},
			expectedError:  nil,
			expectedResult: testResp,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockQuery()

			res, err := mockAuthRepoDb.FindByEmail(tt.inputEmail)

			assert.Equal(t, tt.expectedResult, res)
			assert.ErrorIs(t, err, tt.expectedError)

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
