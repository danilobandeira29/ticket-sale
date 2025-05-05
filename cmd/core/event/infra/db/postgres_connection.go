package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"sync"
)

var (
	once sync.Once
	db   *sql.DB
	errC error
)

func PostgresConn() (*sql.DB, error) {
	once.Do(func() {
		db, errC = sql.Open("postgres", "postgres://postgres:1234@app_db:5432/mba_ddd?sslmode=disable")
		if errC != nil {
			return
		}
	})
	return db, errC
}
