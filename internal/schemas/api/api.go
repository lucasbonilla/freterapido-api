package api

type Config struct {
	APIPort string
	RunType string
}

func NewConfig(APIPort string, RunType string) *Config {
	return &Config{
		APIPort: APIPort,
		RunType: RunType,
	}
}
