package database

import (
	"database/sql"
)

func ExecSQLStatement(tableToCreate string, db *sql.DB) error {
	_, err := db.Exec(tableToCreate)
	if err != nil {
		return err
	}
	return nil
}

func ExecSQLDeleteStatement(tableToCreate string, idToDelete string, db *sql.DB) error {
	_, err := db.Exec(tableToCreate, idToDelete)
	if err != nil {
		return err
	}
	return nil
}