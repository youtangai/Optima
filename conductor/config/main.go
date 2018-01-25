package config

import "github.com/kelseyhightower/envconfig"

type Configuration struct {
	DBUser   string `envconfig:"OPTIMA_DB_USER" default:"zun"`
	DBHost   string `envconfig:"OPTIMA_DB_HOST" default:"localhost"`
	DBPort   string `envconfig:"OPTIMA_DB_PORT" default:"3306"`
	DBName   string `envconfig:"OPTIMA_DB_NAME" default:"zun"`
	DBPasswd string `envconfig:"OPTIMA_DB_PASSWD" default:"199507620"`
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

func DBUser() string {
	return config.DBUser
}

func DBHost() string {
	return config.DBHost
}

func DBPort() string {
	return config.DBPort
}

func DBName() string {
	return config.DBName
}

func DBPasswd() string {
	return config.DBPasswd
}
