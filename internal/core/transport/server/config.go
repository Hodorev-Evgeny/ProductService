package server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	ADDR     int           `envconfig:"ADDR"`
	TIMEZONE time.Location `envconfig:"TIMEZONE"`
	TIMEOUT  time.Duration `envconfig:"TIMEOUT"`
}

func getServerConfig() (ServerConfig, error) {
	var cfg ServerConfig
	if err := envconfig.Process("SERVER", &cfg); err != nil {
		return ServerConfig{}, fmt.Errorf("could not process server config: %w", err)
	}

	return cfg, nil
}

func MustGetServerConfig() ServerConfig {
	cfg, err := getServerConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
