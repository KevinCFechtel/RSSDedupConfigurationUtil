package Configuration

type Configuration struct {
	DatabaseServer            string            `json:"DatabaseServer"`
	DatabasePort              int               `json:"DatabasePort"`
	DatabaseName              string            `json:"DatabaseName"`
	DatabaseUser              string            `json:"DatabaseUser"`
	DatabasePassword          string            `json:"DatabasePassword"`
	SSLMode                   string            `json:"sslmode"`
	HttpServerPort            string            `json:"HttpServerPort"`
	Err                       string
}

func CreateNewConfiguration() Configuration {
	config := Configuration{}

	return config
}