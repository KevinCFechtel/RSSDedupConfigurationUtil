package database

import (
	"database/sql"

	DBModels "github.com/KevinCFechtel/RSSDedupConfigurationUtil/models/database_models"
)

func CheckAndCreateTables(db *sql.DB) error {
	checkTableRssFeedItemIDS := `
	CREATE TABLE  IF NOT EXISTS rssFeedItemIDS (
		id SERIAL PRIMARY KEY, 
		item_id VARCHAR (50) UNIQUE NOT NULL, 
		feed VARCHAR (50) NOT NULL
		);
	`

	checkTableRssDedupConfig := `
	CREATE TABLE  IF NOT EXISTS rssDedupConfig (
		id SERIAL PRIMARY KEY, 
		httpEndpoint TEXT NOT NULL, 
		feedName TEXT NOT NULL,
		feedURL TEXT NOT NULL,
		feedIDFromStartOrEnd TEXT NOT NULL,
		feedIDLength INTEGER NOT NULL,
		feedIDFromStartOrEndLength INTEGER NOT NULL,
		feedIconURL TEXT NOT NULL,
		artikelImageTag TEXT NOT NULL
		);
	`
	_, err := db.Exec(checkTableRssFeedItemIDS)
	if err != nil {
		return err
	}

	_, err = db.Exec(checkTableRssDedupConfig)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFromRssDedupConfig(idToDelete string, db *sql.DB) error {
	deleteStatemen := "DELETE FROM rssDedupConfig WHERE id = $1"
	_, err := db.Exec(deleteStatemen, idToDelete)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFromRssFeedItemIds(feedToDelete string, db *sql.DB) error {
	deleteStatemen := "DELETE FROM rssFeedItemIDS WHERE feed = $1"
	_, err := db.Exec(deleteStatemen, feedToDelete)
	if err != nil {
		return err
	}
	return nil
}

func InsertNewConfig(configuration DBModels.RssDedupConfig, db *sql.DB) error {

	insertstmt := `
		INSERT INTO rssDedupConfig (httpEndpoint, 
									feedName,
									feedURL,
									feedIDFromStartOrEnd,
									feedIDLength,
									feedIDFromStartOrEndLength,
									feedIconURL,
									artikelImageTag) 
		values($1, $2, $3, $4, $5, $6, $7, $8)`

    _, err := db.Exec(insertstmt, configuration.HttpEndpoint, configuration.FeedName, configuration.FeedURL, configuration.FeedIDFromStartOrEnd, configuration.FeedIDLength, configuration.FeedIDFromStartOrEndLength, configuration.FeedIconURL, configuration.ArtikelImageTag)
	if err != nil {
		return err
	}
	return nil
}

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

func ReadRssDedupConfigItem(idToSelect string, db *sql.DB) (DBModels.RssDedupConfig, error) {
	item := DBModels.CreateNewRssDedupConfig()
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
		WHERE id = $1`, idToSelect)
	if err != nil {
		return item, err
	}
	defer rows.Close()
	for rows.Next() {
		
		err := rows.Scan(&item.Id, &item.HttpEndpoint, &item.FeedName, &item.FeedURL, &item.FeedIDFromStartOrEnd, &item.FeedIDLength, &item.FeedIDFromStartOrEndLength, &item.FeedIconURL, &item.ArtikelImageTag)
		if err != nil {
			return item, err
		}
	}
	err = rows.Err()
	if err != nil {
		return item, err
	}
	return item, nil
}

func EditRssDedupConfigItem(itemToEdit DBModels.RssDedupConfig, db *sql.DB) error {
	_, err := db.Exec(`
		UPDATE rssDedupConfig SET httpEndpoint = $1, 
			feedName = $2,
			feedURL = $3,
			feedIDFromStartOrEnd = $4,
			feedIDLength = $5,
			feedIDFromStartOrEndLength = $6,
			feedIconURL = $7,
			artikelImageTag = $8
		WHERE id = $9`, itemToEdit.HttpEndpoint, itemToEdit.FeedName, itemToEdit.FeedURL, itemToEdit.FeedIDFromStartOrEnd, itemToEdit.FeedIDLength, itemToEdit.FeedIDFromStartOrEndLength, itemToEdit.FeedIconURL, itemToEdit.ArtikelImageTag, itemToEdit.Id)
	if err != nil {
		return err
	}
	return nil
}