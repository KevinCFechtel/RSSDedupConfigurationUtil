package database

import (
	"database/sql"

	DBModels "github.com/KevinCFechtel/RSSDedupConfigurationUtil/models/database_models"
)

func ReadTableContentOfRssDedupConfig(db *sql.DB) ([]DBModels.RssDedupConfig, error) {
	var returnModels []DBModels.RssDedupConfig
	rows, err := db.Query(`
		SELECT id, 
			httpEndpoint, 
			feedName,
			feedURL,
			feedIDFromStartOrEnd,
			feedIDLength,
			feedIDFromStartOrEndLength,
			feedIconURL,
			artikelImageTag
		FROM rssDedupConfig 
		ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		newRow := DBModels.CreateNewRssDedupConfig()
		err := rows.Scan(&newRow.Id, &newRow.HttpEndpoint, &newRow.FeedName, &newRow.FeedURL, &newRow.FeedIDFromStartOrEnd, &newRow.FeedIDLength, &newRow.FeedIDFromStartOrEndLength, &newRow.FeedIconURL, &newRow.ArtikelImageTag)
		if err != nil {
			return nil, err
		}
		returnModels = append(returnModels, newRow)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return returnModels, nil
}