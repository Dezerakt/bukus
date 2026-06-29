package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"time"
)

type (
	Config struct {
		App     app
		HTTP    http
		Log     log
		PG      pg
		JWT     jwt
		Metrics metrics
		Swagger swagger
	}

	// App -.
	app struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	// HTTP -.
	http struct {
		Port           string `env:"HTTP_PORT,required"`
		UsePreforkMode bool   `env:"HTTP_USE_PREFORK_MODE" envDefault:"false"`
	}

	// Log -.
	log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	// PG -.
	pg struct {
		PoolMax  int    `env:"PG_POOL_MAX,required"`
		Host     string `env:"PG_HOST,required"`
		User     string `env:"PG_USER,required"`
		DB       string `env:"PG_DB,required"`
		Port     string `env:"PG_PORT,required"`
		Password string `env:"PG_PASSWORD,required"`
	}

	// JWT -.
	jwt struct {
		Secret      string        `env:"JWT_SECRET,required"`
		TokenExpiry time.Duration `env:"JWT_TOKEN_EXPIRY" envDefault:"24h"`
	}

	// Metrics -.
	metrics struct {
		Enabled bool `env:"METRICS_ENABLED" envDefault:"true"`
	}

	// Swagger -.
	swagger struct {
		Enabled bool `env:"SWAGGER_ENABLED" envDefault:"false"`
	}
)

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
