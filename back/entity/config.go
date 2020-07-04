package entity

import "fmt"

type Config struct {
	// General
	IncipitHost string `env:"INCIPIT_HOST"`
	IncipitPort int    `env:"INCIPIT_PORT"`

	// MySQL
	MySQLHost     string `env:"INCIPIT_MYSQL_HOST"`
	MySQLPort     int    `env:"INCIPIT_MYSQL_PORT"`
	MySQLUser     string `env:"INCIPIT_MYSQL_USER"`
	MySQLPassword string `env:"INCIPIT_MYSQL_PASS"`
}

func (c *Config) GetDSN() (dsn string) {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/incipit?charset=utf8mb4&parseTime=true", c.MySQLUser, c.MySQLPassword, c.MySQLHost, c.MySQLPort)
}
