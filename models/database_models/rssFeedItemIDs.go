package database_models

type RssFeedItemIDs struct {
	Id                    int
	Item_id               string
	Feed                  string
}

func CreateNewRssFeedItemID() RssFeedItemIDs {
	config := RssFeedItemIDs{}

	return config
}