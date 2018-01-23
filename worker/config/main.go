package config

import "github.com/kelseyhightower/envconfig"

type Configuration struct {
	Port string `envconfig:"PORT" default:"62071"`
}

var (
	config Configuration
)

const (
	prefix = "TBS"
)

func init() {
	envconfig.MustProcess(prefix, &config)
}

func reload() {
	envconfig.Process(prefix, &config)
}

func Port() string {
	return config.Port
}
