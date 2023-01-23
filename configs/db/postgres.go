package configs

import (
	"fmt"
	"os"
	"strconv"
)

// Postgres config object
type PostgresConfig struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

// GetPostgresConnectionInfo returns postgres connection uri
func (c PostgresConfig) GetPostgresConnectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Name)
	}

	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name)
}

// GetPostgresConfig returns Postgres config object
func GetPostgresConfig() PostgresConfig {
	port, portErr := strconv.Atoi(os.Getenv("DB_PORT"))
	if portErr != nil {
		panic(portErr)
	}
	return PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
}
