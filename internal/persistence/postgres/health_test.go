package postgres

import "testing"

func TestHealthRepo_Ping(t *testing.T) {
	db, mock := createConn(t)
	hr := NewHealthRepo(db)

	mock.ExpectPing()

	if err := hr.Ping(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
