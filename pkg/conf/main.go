package conf

type Config struct {
	BaseUrl string `json:"baseUrl"`
}

func NewConfig() Config {
	return Config{
		BaseUrl: "https://",
	}
}
