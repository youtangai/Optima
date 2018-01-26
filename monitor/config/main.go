package config

type Configuration struct {
	ConductorHost string
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
