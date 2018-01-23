package config

import "github.com/kelseyhightower/envconfig"

type Configuration struct {
	DBUser   string `envconfig:"DB_USER" default:"zun"`
	DBHost   string `envconfig:"DB_HOST" default:"localhost"`
	DBPort   string `envconfig:"DB_PORT" default:"3306"`
	DBName   string `envconfig:"DB_NAME" default:"zun-optima"`
	DBPasswd string `envconfig:"DB_PASSWD" default:"199507620"`
	DBMS     string `envconfig:"DBMS" default:"mysql"`
	GinPort  string `envconfig:"GIN_PORT" default:"62070"`
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

func DBMS() string {
	return config.DBMS
}

func GinPort() string {
	return config.GinPort
}
