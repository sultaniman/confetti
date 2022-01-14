package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Connect godoc
// Creates connections to database
func Connect(dbUri string) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", dbUri)
}
