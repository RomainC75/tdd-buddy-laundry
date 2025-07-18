package config

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Name     string
}

var config *Config
