package config

import (
	"fmt"
	"os"
)

var Cfg *Config

type Config struct {
	RxCfg  ReindexerConfig
	Server Server
}

var (
	NoEnvVarsError = fmt.Errorf("no one environmental variables were found")
)

type ReindexerConfig struct {
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	Namespaces Namespaces
}

type Server struct {
	Host string
	Port string
}

type Namespaces struct {
	Teachers string
	Groups   string
	Names    string
}

// TODO wrap by logs
func Init() error {
	cfg := &Config{
		RxCfg: ReindexerConfig{
			Host:     os.Getenv("RX_HOST"),
			Port:     os.Getenv("RX_PORT"),
			Username: os.Getenv("RX_USERNAME"),
			Password: os.Getenv("RX_PASSWORD"),
			Database: os.Getenv("RX_DATABASE"),
			Namespaces: Namespaces{
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
