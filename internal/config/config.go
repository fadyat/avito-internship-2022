package config

import "fmt"

type HTTPConfig struct {
	Debug          bool   `envconfig:"DEBUG" default:"false"`
	HTTPPort       string `envconfig:"HTTP_PORT" default:"80"`
	DatabaseDriver string `envconfig:"DB_DRIVER" default:"postgres"`
	DatabaseHost   string `envconfig:"DB_HOST" default:"localhost"`
	DatabasePort   string `envconfig:"DB_PORT" default:"5432"`
	DatabaseUser   string `envconfig:"DB_USER" default:"postgres"`
	DatabasePass   string `envconfig:"DB_PASS" default:"postgres"`
	DatabaseName   string `envconfig:"DB_NAME" default:"postgres"`
	DatabaseSSL    string `envconfig:"DB_SSL" default:"disable"`
}

func (cfg *HTTPConfig) GetConnectionString() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DatabaseDriver,
		cfg.DatabaseUser,
		cfg.DatabasePass,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName,
		cfg.DatabaseSSL,
	)
}
