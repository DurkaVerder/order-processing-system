package postgres

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateStatus_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error create mockdb: %v", err)
	}
	defer db.Close()

	p := NewPostgres(db)

	mock.ExpectExec("UPDATE orders SET status = \\$1 WHERE id = \\$2").
		WithArgs("in-processing", 1).WillReturnResult(sqlmock.NewResult(0, 1))

	err = p.UpdateStatus(1, "in-processing")

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateStatus_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error create mockdb: %v", err)
	}
	defer db.Close()

	p := NewPostgres(db)

	expectedError := errors.New("error connection db")
	mock.ExpectExec("UPDATE orders SET status = \\$1 WHERE id = \\$2").
		WithArgs("in-processing", 1).WillReturnError(expectedError)

	err = p.UpdateStatus(1, "in-processing")

	assert.EqualError(t, err, expectedError.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateRecordStatus_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error create mockdb: %v", err)
	}
	defer db.Close()

	p := NewPostgres(db)

	mock.ExpectExec("INSERT INTO order_status_history \\(order_id, status\\) VALUES \\(\\$1, \\$2\\)").
		WithArgs(1, "in-processing").WillReturnResult(sqlmock.NewResult(0, 1))

	err = p.CreateRecordStatus(1, "in-processing")

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateRecordStatus_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error create mockdb: %v", err)
	}
	defer db.Close()

	p := NewPostgres(db)

	expectedError := errors.New("error connection db")
	mock.ExpectExec("INSERT INTO order_status_history \\(order_id, status\\) VALUES \\(\\$1, \\$2\\)").
		WithArgs(1, "in-processing").WillReturnError(expectedError)

	err = p.CreateRecordStatus(1, "in-processing")

	assert.EqualError(t, err, expectedError.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}
