package database

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("pgx", "postgresql://postgres:Ahyaeki12@localhost:5432/mydb")
	if err != nil {
		return nil, err
	}
	return db, nil
}
