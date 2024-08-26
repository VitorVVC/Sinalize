package models

type Config struct {
	API APIConfig
	DB  DBConfig
	DSN DSNService
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

type DSNService struct {
	Dsn string
}
