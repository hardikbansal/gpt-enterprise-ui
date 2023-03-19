package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	// cache
	CacheLocation string `env:"CACHE_LOCATION"`

	Port    string `env:"PORT"`
	IsDebug bool   `env:"IS_DEBUG"`

	// // default database
	// DefaultDatabaseName     string `env:"DEFAULT_DATABASE_NAME"`
	// DefaultDatabasePassword string `env:"DEFAULT_DATABASE_PASSWORD"`
	// DefaultDatabaseUser     string `env:"DEFAULT_DATABASE_USER"`
	// DefaultDatabaseHost     string `env:"DEFAULT_DATABASE_HOST"`
	// DefaultDatabasePort     uint16 `env:"DEFAULT_DATABASE_PORT"`

	// // read database
	// ReadDatabaseName     string `env:"READ_DATABASE_NAME"`
	// ReadDatabaseUser     string `env:"READ_DATABASE_USER"`
	// ReadDatabasePassword string `env:"READ_DATABASE_PASSWORD"`
	// ReadDatabaseHost     string `env:"READ_DATABASE_HOST"`
	// ReadDatabasePort     uint16 `env:"READ_DATABASE_PORT"`

	// // write database
	// WriteDatabaseName     string `env:"WRITE_DATABASE_NAME"`
	// WriteDatabaseUser     string `env:"WRITE_DATABASE_USER"`
	// WriteDatabasePassword string `env:"WRITE_DATABASE_PASSWORD"`
	// WriteDatabaseHost     string `env:"WRITE_DATABASE_HOST"`
	// WriteDatabasePort     uint16 `env:"WRITE_DATABASE_PORT"`

	IsSentryActivated    bool    `env:"ACTIVATE_SENTRY"`
	SentryKey            string  `env:"SENTRY_KEY"`
	SentryUrl            string  `env:"SENTRY_URL"`
	SentryEnvironment    string  `env:"SENTRY_ENVIRONMENT"`
	SentryLevelCode      string  `env:"SENTRY_LEVEL_CODE"`
	SentryEventLevelCode string  `env:"SENTRY_EVENT_LEVEL_CODE"`
	SentrySampleRatePct  float64 `env:"SENTRY_API_PERF_SAMPLE_RATE_PCT"`

	GitCommitId string `env:"GIT_COMMIT"`
}

var lock = &sync.Mutex{}

var appConfig *AppConfig

func loadConfig() (config AppConfig, err error) {
	configPath := "./config/app.env"
	err = godotenv.Load(configPath)
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	if err = env.Parse(&config); err != nil {
		log.Printf("%+v\n", err)
	}
	return
}

func GetInstance() *AppConfig {
	if appConfig == nil {
		lock.Lock()
		defer lock.Unlock()
		if appConfig == nil {
			log.Print("Creating single instance now.")
			config, err := loadConfig()
			appConfig = &config
			if err != nil {
				panic(err)
			}
		}
	}
	return appConfig
}
