package database

import (
	"database/sql"

	DBModels "github.com/KevinCFechtel/RSSDedupConfigurationUtil/models/database_models"
)

func InsertNewConfig(configuration DBModels.RssDedupConfig, db *sql.DB) error {

	insertstmt := `
		INSERT INTO rssDedupConfig (httpEndpoint, 
									feedName,
									feedURL,
									feedIDFromStartOrEnd,
									feedIDLength,
									feedIDFromStartOrEndLength,
									feedIconURL) 
		values($1, $2, $3, $4, $5, $6, $7)`

    _, err := db.Exec(insertstmt, configuration.HttpEndpoint, configuration.FeedName, configuration.FeedURL, configuration.FeedIDFromStartOrEnd, configuration.FeedIDLength, configuration.FeedIDFromStartOrEndLength, configuration.FeedIconURL)
	if err != nil {
		return err
	}
	return nil
}