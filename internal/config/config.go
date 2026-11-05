package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
	DatabaseURL string
}

func MustLoad() Config {
  godotenv.Load()
  port := os.Getenv("PORT")
  if port == "" {
	panic("PORT environment variable is not set")
  }

  env := os.Getenv("ENV")
  if env == "" {
	panic("ENV environment variable is not set")
  }

  databaseURL := os.Getenv("DATABASE_URL")
  if databaseURL == "" {
	panic("DATABASE_URL environment variable is not set")
  }

  return Config{
	Port: port,
	Env:  env,
	DatabaseURL: databaseURL,
  }
}
