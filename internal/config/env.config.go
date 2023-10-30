package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// server environment variables
var (
	StageStatus       = os.Getenv("STAGE_STATUS")
	Host              = os.Getenv("HOST")
	Port              = os.Getenv("PORT")
	ServerReadTimeout = os.Getenv("SERVER_READ_TIMEOUT")
)

// authentication environment variables
var (
	JwtSecret = os.Getenv("JWT_SECRET_KEY")
	JwtCost   = os.Getenv("JWT_COST")
)

// database environment variables
var (
	DbHost     = os.Getenv("DB_HOST")
	DbPort     = os.Getenv("DB_PORT")
	DbUser     = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName     = os.Getenv("DB_NAME")
	DbSSLMode  = os.Getenv("DB_SSL_MODE")
)

// cache environment variables
var (
	RedisHost     = os.Getenv("REDIS_HOST")
	RedisPort     = os.Getenv("REDIS_PORT")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
)

// sentry environment variables
var (
	SentryDsn = os.Getenv("SENTRY_DSN")
)

// email
var (
	Email            = os.Getenv("EMAIL")
	EmailAppPassword = os.Getenv("EMAIL_APP_PASSWORD")
)
