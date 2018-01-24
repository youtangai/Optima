package config

import "github.com/kelseyhightower/envconfig"

type Configuration struct {
	Port          string `envconfig:"PORT" default:"62071"`
	ConductorHost string `envconfig:"CONDUCTOR_HOST" default:"conductor"`
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

func ConductorHost() string {
	return config.ConductorHost
}
