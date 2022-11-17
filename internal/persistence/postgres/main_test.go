package postgres

import (
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"reflect"
	"testing"
)

func createConn(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %v", err)
	}

	return sqlx.NewDb(db, "sqlmock"), mock
}

func compareResults(t *testing.T, exp, got interface{}, mock sqlmock.Sqlmock) {
	t.Helper()

	if !reflect.DeepEqual(exp, got) {
		t.Fatalf("expected result: %v, got: %v", exp, got)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %v", err)
	}
}

func compareErrors(t *testing.T, exp, got error) {
	t.Helper()

	if exp != got {
		t.Fatalf("expected error: %v, got: %v", exp, got)
	}
}

func createErrorMock(t *testing.T, mock sqlmock.Sqlmock, query string, err error, args ...driver.Value) {
	t.Helper()

	mock.ExpectQuery(query).WithArgs(args...).WillReturnError(err)
}

func createSuccessMock(t *testing.T, mock sqlmock.Sqlmock, query string, rows *sqlmock.Rows, args ...driver.Value) {
	t.Helper()

	mock.ExpectQuery(query).WithArgs(args...).WillReturnRows(rows)
}
