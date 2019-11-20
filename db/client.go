package db

import (
	// Import this so we don't have to use qm.Limit etc.
	"database/sql"
	_ "github.com/lib/pq" // here
	"github.com/volatiletech/sqlboiler/boil"
	_ "github.com/volatiletech/sqlboiler/queries/qm"

	"sync"
)

var doOnce sync.Once
var db *sql.DB
var err error

func GetClient() (*sql.DB, error) {
	doOnce.Do(func() {
		db, err = sql.Open("postgres", "dbname=postgres user=postgres password=password1 sslmode=disable")
		boil.SetDB(db)
	})
	return db, err
}
