package config

func Set(username string, password string, host string, port int, dbname string) {
	config = &Config{
		Db: DbConfig{username, password, host, port, dbname},
	}
}
