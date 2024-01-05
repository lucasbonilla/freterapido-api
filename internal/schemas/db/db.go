package db

type Config struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func NewConfig(host string, port string, user string, pass string, database string) *Config {
	return &Config{
		Host:     host,
		Port:     port,
		User:     user,
		Pass:     pass,
		Database: database,
	}
}
