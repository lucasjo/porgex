package config

import "gopkg.in/gcfg.v1"

type Config struct {
	Mongo struct {
		Host         string
		Port         string
		Username     string
		Password     string
		Databasename string
	}
}

func NewConfig(filepath string) (Config, error) {

	var config Config
	err := gcfg.ReadFileInto(&config, filepath)

	return config, err
}
