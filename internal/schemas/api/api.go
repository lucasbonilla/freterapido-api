package api

type Config struct {
	APIPort string
	RunType string
}

func NewConfig(apiPort string, runType string) *Config {
	return &Config{
		APIPort: apiPort,
		RunType: runType,
	}
}
