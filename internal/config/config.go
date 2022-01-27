package config

import (
	"io/ioutil"
	toml "github.com/pelletier/go-toml/v2"
)


type Database struct {
	Username string
	Password string
	Database string
	Port int
	Hostname string
}

type Log struct {
	File string
	Level string
}

type Server struct {
	Url string
	Port int
}

type serviceConfig struct {
	Version int
	Database Database `toml:"database"`
	Log Log `toml:"log"`
	Server Server `toml:"server"`
}


func parseConfig(file string) serviceConfig {
	data, read_err := ioutil.ReadFile(file)

	if read_err != nil {
		panic(read_err)
	}

	var configuration serviceConfig

	config_err := toml.Unmarshal([]byte(data), &configuration)
	if config_err != nil {
		panic(config_err)
	}

	return configuration
}


func ReadDatabaseConfig(file string) Database {
	configuration := parseConfig(file)
	return configuration.Database
}

func ReadLogConfig(file string) Log {
	configuration := parseConfig(file)
	return configuration.Log
}

func ReadServerConfig(file string) Server {
	configuration := parseConfig(file)
	return configuration.Server
}
