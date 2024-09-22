package Configuration

type Configuration struct {
	Server                    string            `json:"server"`
	Port                      int               `json:"port"`
	Database                  string            `json:"database"`
	User                      string            `json:"user"`
	Pass                      string            `json:"pass"`
	SSLMode                 string            `json:"sslmode"`
	Err                       string
}

func CreateNewConfiguration() Configuration {
	config := Configuration{}

	return config
}