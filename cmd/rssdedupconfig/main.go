package main

import (
	"database/sql"
	"fmt"
	"strconv"

	DatabaseFunctions "github.com/KevinCFechtel/RSSDedupConfigurationUtil/functions/database"
	Configuration "github.com/KevinCFechtel/RSSDedupConfigurationUtil/models/configuration"
	DBModels "github.com/KevinCFechtel/RSSDedupConfigurationUtil/models/database_models"
	configHandler "github.com/KevinCFechtel/goConfigHandler"
	_ "github.com/lib/pq"
	"github.com/pterm/pterm"
)

func main() {
	options := []string{"List Table rssFeedItemIDS", "List Table rssDedupConfig", "Add new Configuration", "Exit"}
	pterm.Printfln("Please provide the path to the config file:")
	configFilePath, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue("config.json").Show()
	pterm.Println()

	configuration := Configuration.CreateNewConfiguration()
	err := configHandler.GetConfig("localFile", configFilePath, &configuration, "File not found")
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=%s",
    configuration.Server, configuration.Port, configuration.User, configuration.Pass, configuration.Database, configuration.SSLMode)

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
		feedIconURL TEXT NOT NULL
		);
	`

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	}
	err = db.Ping()
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	}

	err = DatabaseFunctions.CreateTables(checkTableRssFeedItemIDS, db)
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	}

	err = DatabaseFunctions.CreateTables(checkTableRssDedupConfig, db)
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	} 
	selectedOption := ""
	for selectedOption != "Exit" {
		selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(options).Show()
		switch selectedOption {
			case "List Table rssFeedItemIDS": 
				lines, err := DatabaseFunctions.ReadTableContentOfRssFeedItemIDs(db)
				if err != nil {
					pterm.Fatal.PrintOnError(err)
				} 
				selectResult := pterm.TableData{
					{"Id", "Item_id", "Feed"},
				}
		
				for _, sqlLine := range lines {
					newTableLine := []string{strconv.Itoa(sqlLine.Id), sqlLine.Item_id, sqlLine.Feed}
					selectResult = append(selectResult, newTableLine)
				}
			
				pterm.DefaultTable.WithHasHeader().WithData(selectResult).Render()
			case "List Table rssDedupConfig": 
				lines, err := DatabaseFunctions.ReadTableContentOfRssDedupConfig(db)
				if err != nil {
					pterm.Fatal.PrintOnError(err)
				} 
				selectResult := pterm.TableData{
					{"Id", "HttpEndpoint", "FeedName", "FeedURL", "FeedIDFromStartOrEnd", "FeedIDLength", "FeedIDFromStartOrEndLength", "FeedIconURL"},
				}
		
				for _, sqlLine := range lines {
					newTableLine := []string{strconv.Itoa(sqlLine.Id), sqlLine.HttpEndpoint, sqlLine.FeedName, sqlLine.FeedURL, sqlLine.FeedIDFromStartOrEnd, strconv.Itoa(sqlLine.FeedIDLength), strconv.Itoa(sqlLine.FeedIDFromStartOrEndLength), sqlLine.FeedIconURL}
					selectResult = append(selectResult, newTableLine)
				}
		
				pterm.DefaultTable.WithHasHeader().WithData(selectResult).Render()
			case "Add new Configuration": 
				newConfig := DBModels.CreateNewRssDedupConfig()

				pterm.Println("Please enter HttpEndpoint:")
				result, _ := pterm.DefaultInteractiveTextInput.Show()
				newConfig.HttpEndpoint = result

				pterm.Println("Please enter FeedName:")
				result, _ = pterm.DefaultInteractiveTextInput.Show()
				newConfig.FeedName = result

				pterm.Println("Please enter FeedURL:")
				result, _ = pterm.DefaultInteractiveTextInput.Show()
				newConfig.FeedURL = result

				pterm.Println("Please enter FeedIDFromStartOrEnd:")
				result, _ = pterm.DefaultInteractiveTextInput.Show()
				newConfig.FeedIDFromStartOrEnd = result

				pterm.Println("Please enter FeedIDLength:")
				result, _ = pterm.DefaultInteractiveTextInput.Show()
				newString, err := strconv.Atoi(result)
				if err != nil {
					pterm.Fatal.PrintOnError(err)
				} 
				newConfig.FeedIDLength = newString

				pterm.Println("Please enter FeedIDFromStartOrEndLength:")
				result, _ = pterm.DefaultInteractiveTextInput.Show()
				newString, err = strconv.Atoi(result)
				if err != nil {
					pterm.Fatal.PrintOnError(err)
				} 
				newConfig.FeedIDFromStartOrEndLength = newString

				pterm.Println("Please enter FeedIconURL:")
				result, _ = pterm.DefaultInteractiveTextInput.Show()
				newConfig.FeedIconURL = result

				DatabaseFunctions.InsertNewConfig(newConfig, db)
				if err != nil {
					pterm.Fatal.PrintOnError(err)
				} else {
					pterm.Println("Config successful added")
				}
				
		}
	}

	
}