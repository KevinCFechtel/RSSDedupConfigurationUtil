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
	options := []string{"List Table rssFeedItemIDS", "List Table rssDedupConfig", "Add new Configuration", "Edit Configuration", "Delete Configuration", "Exit"}
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
    configuration.DatabaseServer, configuration.DatabasePort, configuration.DatabaseUser, configuration.DatabasePassword, configuration.DatabaseName, configuration.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	}
	err = db.Ping()
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	}

	err = DatabaseFunctions.CheckAndCreateTables(db)
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	}

	selectedOption := ""
	for selectedOption != "Exit" {
		selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(options).Show()
		switch selectedOption {
			case "List Table rssFeedItemIDS": 
				listRssFeedItemIdsTable(db)
			case "List Table rssDedupConfig": 
				listRssDedupConfigTable(db)
			case "Add new Configuration": 
				addConfiguration(db)
			case "Edit Configuration":
				editConfiguration(db)
			case "Delete Configuration":
				deleteConfiguration(db)	
		}
	}

	
}

func listRssFeedItemIdsTable(db *sql.DB) {
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
}

func listRssDedupConfigTable(db *sql.DB) {
	lines, err := DatabaseFunctions.ReadTableContentOfRssDedupConfig(db)
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	}
	selectResult := pterm.TableData{
		{"Id", "HttpEndpoint", "FeedName", "FeedURL", "FeedIDFromStartOrEnd", "FeedIDLength", "FeedIDFromStartOrEndLength", "FeedIconURL", "artikelImageTag"},
	}

	for _, sqlLine := range lines {
		newTableLine := []string{strconv.Itoa(sqlLine.Id), sqlLine.HttpEndpoint, sqlLine.FeedName, sqlLine.FeedURL, sqlLine.FeedIDFromStartOrEnd, strconv.Itoa(sqlLine.FeedIDLength), strconv.Itoa(sqlLine.FeedIDFromStartOrEndLength), sqlLine.FeedIconURL, sqlLine.ArtikelImageTag}
		selectResult = append(selectResult, newTableLine)
	}

	pterm.DefaultTable.WithHasHeader().WithData(selectResult).Render()
}

func addConfiguration(db *sql.DB) {
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

	pterm.Println("Please enter ArtikelImageTag:")
	result, _ = pterm.DefaultInteractiveTextInput.Show()
	newConfig.ArtikelImageTag = result

	err = DatabaseFunctions.InsertNewConfig(newConfig, db)
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	} else {
		pterm.Println("Config successful added")
	}
}

func deleteConfiguration(db *sql.DB) {
	pterm.Println("Please enter the id to delete:")
	result, _ := pterm.DefaultInteractiveTextInput.Show()

	configItemToDelete, err := DatabaseFunctions.ReadRssDedupConfigItem(result, db)
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	}

	err = DatabaseFunctions.DeleteFromRssFeedItemIds(configItemToDelete.FeedName, db)
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	}

	DatabaseFunctions.DeleteFromRssDedupConfig(result, db)
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	} else {
		pterm.Println("Config successful deleted")
	}
}

func editConfiguration(db *sql.DB) {
	pterm.Println("Please enter the id to edit:")
	result, _ := pterm.DefaultInteractiveTextInput.Show()

	configItemToEdit, err := DatabaseFunctions.ReadRssDedupConfigItem(result, db)
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	}

	pterm.Println("Please enter HttpEndpoint:")
	result, _ = pterm.DefaultInteractiveTextInput.WithDefaultValue(configItemToEdit.HttpEndpoint).Show()
	configItemToEdit.HttpEndpoint = result

	pterm.Println("Please enter FeedName:")
	result, _ = pterm.DefaultInteractiveTextInput.WithDefaultValue(configItemToEdit.FeedName).Show()
	configItemToEdit.FeedName = result

	pterm.Println("Please enter FeedURL:")
	result, _ = pterm.DefaultInteractiveTextInput.WithDefaultValue(configItemToEdit.FeedURL).Show()
	configItemToEdit.FeedURL = result

	pterm.Println("Please enter FeedIDFromStartOrEnd:")
	result, _ = pterm.DefaultInteractiveTextInput.WithDefaultValue(configItemToEdit.FeedIDFromStartOrEnd).Show()
	configItemToEdit.FeedIDFromStartOrEnd = result

	pterm.Println("Please enter FeedIDLength:")
	result, _ = pterm.DefaultInteractiveTextInput.WithDefaultValue(strconv.Itoa(configItemToEdit.FeedIDLength)).Show()
	newInt, err := strconv.Atoi(result)
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	}
	configItemToEdit.FeedIDLength = newInt

	pterm.Println("Please enter FeedIDFromStartOrEndLength:")
	result, _ = pterm.DefaultInteractiveTextInput.WithDefaultValue(strconv.Itoa(configItemToEdit.FeedIDFromStartOrEndLength)).Show()
	newInt, err = strconv.Atoi(result)
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	}
	configItemToEdit.FeedIDFromStartOrEndLength = newInt

	pterm.Println("Please enter FeedIconURL:")
	result, _ = pterm.DefaultInteractiveTextInput.WithDefaultValue(configItemToEdit.FeedIconURL).Show()
	configItemToEdit.FeedIconURL = result

	pterm.Println("Please enter ArtikelImageTag:")
	result, _ = pterm.DefaultInteractiveTextInput.WithDefaultValue(configItemToEdit.ArtikelImageTag).Show()
	configItemToEdit.ArtikelImageTag = result

	err = DatabaseFunctions.EditRssDedupConfigItem(configItemToEdit, db)
	if err != nil {
		pterm.Fatal.PrintOnError(err)
	} else {
		pterm.Println("Config successful edited")
	}
}