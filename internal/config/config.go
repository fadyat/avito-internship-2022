package config

import "fmt"

type HTTPConfig struct {
	Debug          bool   `envconfig:"DEBUG" default:"false" `       //validate:"oneof=true false"`
	HTTPPort       string `envconfig:"HTTP_PORT" default:"80"`       //validate:"required,min=1,max=65535"`
	DatabaseDriver string `envconfig:"DB_DRIVER" default:"postgres"` // validate:"required"`
	DatabaseHost   string `envconfig:"DB_HOST" default:"localhost"`  // validate:"required"`
	DatabasePort   string `envconfig:"DB_PORT" default:"5432"`       // validate:"required,min=1,max=65535"`
	DatabaseUser   string `envconfig:"DB_USER" default:"postgres"`   /// validate:"required"`
	DatabasePass   string `envconfig:"DB_PASS" default:"postgres"`   // validate:"required"`
	DatabaseName   string `envconfig:"DB_NAME" default:"postgres"`   // validate:"required"`
	DatabaseSSL    string `envconfig:"DB_SSL" default:"disable"`     // validate:"oneof=disable require prefer verify-ca verify-full"`
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
