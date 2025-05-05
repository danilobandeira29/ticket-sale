package db

import (
	"context"
	"database/sql"
	"errors"
	"testing"
)

func TestPartnerRepository_FindAll(t *testing.T) {
	database, err := PostgresConn()
	if err != nil {
		t.Errorf("error conn: %v", err)
		return
	}
	repo := NewRepository(database)
	_, err = repo.FindAll(context.Background())
	if errors.Is(err, sql.ErrNoRows) {
		t.Errorf("no rows")
		return
	}
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
}
