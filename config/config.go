package config

import (
	"fmt"
	"github.com/anCreny/IsuctSchedule-Packages/structs"
	"os"
)

var Cfg *Config

type Config struct {
	RxCfg  structs.ReindexerConfig
	Server Server
}

var (
	NoEnvVarsError = fmt.Errorf("no one environmental variables were found")
)

type Server struct {
	Host string
	Port string
}

func Init() error {
	cfg := &Config{
		RxCfg: structs.ReindexerConfig{
			Host:     os.Getenv("RX_HOST"),
			Port:     os.Getenv("RX_PORT"),
			Username: os.Getenv("RX_USERNAME"),
			Password: os.Getenv("RX_PASSWORD"),
			Database: os.Getenv("RX_DATABASE"),
			Namespaces: structs.Namespaces{
				Teachers: os.Getenv("NM_TEACHERS"),
				Groups:   os.Getenv("NM_GROUPS"),
				Names:    os.Getenv("NM_NAMES"),
			},
		},
		Server: Server{
			Host: os.Getenv("CACHE_HOST"),
			Port: os.Getenv("CACHE_PORT"),
		},
	}

	if cSum := cfg.RxCfg.Password +
		cfg.RxCfg.Port +
		cfg.RxCfg.Database +
		cfg.RxCfg.Host +
		cfg.RxCfg.Username +
		cfg.RxCfg.Namespaces.Names +
		cfg.RxCfg.Namespaces.Teachers +
		cfg.RxCfg.Namespaces.Groups; cSum == "" {
		return NoEnvVarsError
	}

	Cfg = cfg

	return nil
}
