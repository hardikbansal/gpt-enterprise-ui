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

	Port    string `env:"API_PORT"`
	IsDebug bool   `env:"IS_DEBUG"`

	EmailDomain string `env:"EMAIL_DOMAIN"`

	ApiKey string `env:"API_KEY"`

	// // default database
	DatabaseName     string `env:"DATABASE_NAME"`
	DatabasePassword string `env:"DATABASE_PASSWORD"`
	DatabaseUser     string `env:"DATABASE_USER"`
	DatabaseHost     string `env:"DATABASE_HOST"`
	DatabasePort     string `env:"DATABASE_PORT"`

	IsSentryActivated bool   `env:"ACTIVATE_SENTRY"`
	SentryKey         string `env:"SENTRY_KEY"`
	SentryUrl         string `env:"SENTRY_URL"`
	SentryEnvironment string `env:"SENTRY_ENVIRONMENT"`
}

var lock = &sync.Mutex{}

var appConfig *AppConfig

func loadConfig() (config AppConfig, err error) {
	configPath := "/config/.env"
	err = godotenv.Load(configPath)
	if err != nil {
		log.Fatalf("unable to load .env file: %e from %s", err, configPath)
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
