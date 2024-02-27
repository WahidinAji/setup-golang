package main

import (
	"os"

	"github.com/rs/zerolog/log"
)

type Config struct {
	Sentry struct {
		DSN string
	}
	Database struct {
		Name     string
		User     string
		Password string
		Host     string
		Port     string
	}
	Environment    string
	Release        string
	Port           string
	SecretPassword string
	Cors           struct {
		AllowOrigins []string
		AllowMethods []string
		AllowHeaders []string
	}
}

type DefaultSettings struct {
	Setting string
	Default bool
}

func getConfig() (config Config) {
	var defaultSettings []DefaultSettings
	dsn, ok := os.LookupEnv("SENTRY_DSN")
	if !ok {
		dsn = ""
		defaultSettings = append(defaultSettings, DefaultSettings{"SENTRY_DSN", true})
	}
	DbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		DbName = "postgres"
		defaultSettings = append(defaultSettings, DefaultSettings{"DB_NAME", true})
	}
	DbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		DbUser = "postgres"
		defaultSettings = append(defaultSettings, DefaultSettings{"DB_USER", true})
	}
	DbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		DbPassword = "password"
		defaultSettings = append(defaultSettings, DefaultSettings{"DB_PASSWORD", true})
	}
	DbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		DbHost = "localhost"
		defaultSettings = append(defaultSettings, DefaultSettings{"DB_HOST", true})
	}
	DbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		DbPort = "5432"
		defaultSettings = append(defaultSettings, DefaultSettings{"DB_PORT", true})
	}
	mode, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		mode = "development"
		defaultSettings = append(defaultSettings, DefaultSettings{"ENVIRONMENT", true})
	}
	release, ok := os.LookupEnv("RELEASE")
	if !ok {
		release = "api-undangan@0.0.1"
		defaultSettings = append(defaultSettings, DefaultSettings{"RELEASE", true})
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
		defaultSettings = append(defaultSettings, DefaultSettings{"PORT", true})
	}

	secretPassword, ok := os.LookupEnv("PASSWORD_SECRET")
	if !ok {
		secretPassword = "secret"
	}

	corsAllowOrigins, ok := os.LookupEnv("CORS_ALLOW_ORIGINS")
	if !ok {
		corsAllowOrigins = "*"
		defaultSettings = append(defaultSettings, DefaultSettings{"CORS_ALLOW_ORIGINS", true})
	}

	corsAllowMethods, ok := os.LookupEnv("CORS_ALLOW_METHODS")
	if !ok {
		corsAllowMethods = "GET,POST"
		defaultSettings = append(defaultSettings, DefaultSettings{"CORS_ALLOW_METHODS", true})
	}

	corsAllowHeaders, ok := os.LookupEnv("CORS_ALLOW_HEADERS")
	if !ok {
		//only allow json
		corsAllowHeaders = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
	}

	if len(defaultSettings) > 0 {
		log.Warn().Msgf("Using default settings for: %+v", defaultSettings)
	}

	config.Sentry.DSN = dsn
	config.Database.Name = DbName
	config.Database.User = DbUser
	config.Database.Password = DbPassword
	config.Database.Host = DbHost
	config.Database.Port = DbPort
	config.Environment = mode
	config.Release = release
	config.Port = port
	config.SecretPassword = secretPassword
	config.Cors.AllowOrigins = []string{corsAllowOrigins}
	config.Cors.AllowMethods = []string{corsAllowMethods}
	config.Cors.AllowHeaders = []string{corsAllowHeaders}
	return
}
