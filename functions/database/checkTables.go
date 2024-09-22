package database

import (
	"database/sql"
)

func CreateTables(tableToCreate string, db *sql.DB) error {
	_, err := db.Exec(tableToCreate)
	if err != nil {
		return err
	}
	return nil
}