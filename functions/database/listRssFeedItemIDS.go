package database

import (
	"database/sql"

	DBModels "github.com/KevinCFechtel/RSSDedupConfigurationUtil/models/database_models"
)

func ReadTableContentOfRssFeedItemIDs(db *sql.DB) ([]DBModels.RssFeedItemIDs, error) {
	var returnModels []DBModels.RssFeedItemIDs
	rows, err := db.Query("SELECT id, item_id, feed FROM rssFeedItemIDS order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		newRow := DBModels.CreateNewRssFeedItemID()
		err := rows.Scan(&newRow.Id, &newRow.Item_id, &newRow.Feed)
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