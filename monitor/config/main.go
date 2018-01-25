package config

type Configuration struct {
	ConductorHost string `envconfig:"CONDUCTOR_HOST" default:"conductor"`
}

var (
	config Configuration
)

func GetConductorHost() string {
	return config.ConductorHost
}

func SetConductorHost(host string) {
	config.ConductorHost = host
}
