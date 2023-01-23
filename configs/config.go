package configs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	dbcfg "app/configs/db"
	cnst "app/constants"
)

// SetupConfigs sets up app configurations
// and returns Config object
func SetupConfigs() Config {

	// Load from dotenv
	LoadDotenv()

	// return configurations
	return GetConfig()
}

// Config object
type Config struct {
	Env      string               `env:"ENV"`
	Host     string               `env:"APP_HOST"`
	Port     string               `env:"APP_PORT"`
	Postgres dbcfg.PostgresConfig `json:"postgres"`

	JWTAccessSecret  string `env:"JWT_ACCESS_SIGN_KEY"`
	JWTRefreshSecret string `env:"JWT_REFRESH_SIGN_KEY"`
	JWTIssuer        string `env:"JWT_ISSUER"`
}

// IsProd Checks if env is production
func (c Config) IsProd() bool {
	return c.Env == cnst.PROD
}

// LoadConfig gets config from .env
func LoadDotenv() {

	log.Print("Loading ENV vars from .env ...")
	currentPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	environmentPath := filepath.Join(currentPath, ".env")

	if err := godotenv.Load(environmentPath); err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
	}
	log.Print("Completed loading!")
}

// GetConfig gets all config for the application
func GetConfig() Config {
	return Config{
		Env:              os.Getenv("ENV"),
		Host:             os.Getenv("APP_HOST"),
		Port:             os.Getenv("APP_PORT"),
		Postgres:         dbcfg.GetPostgresConfig(),
		JWTAccessSecret:  os.Getenv("JWT_ACCESS_SIGN_KEY"),
		JWTRefreshSecret: os.Getenv("JWT_REFRESH_SIGN_KEY"),
		JWTIssuer:        os.Getenv("JWT_ISSUER"),
	}
}
