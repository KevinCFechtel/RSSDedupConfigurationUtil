package database_models

type RssDedupConfig struct {
	Id                         	int
	HttpEndpoint               	string
	FeedName                   	string
	FeedURL                    	string
	FeedIDFromStartOrEnd       	string
	FeedIDLength               	int
	FeedIDFromStartOrEndLength 	int
	FeedIconURL 			    string
}

func CreateNewRssDedupConfig() RssDedupConfig {
	config := RssDedupConfig{}

	return config
}