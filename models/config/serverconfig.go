package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ServerConfig struct {
	Port         int      `yaml:"port"`
	Address      string   `yaml:"address"`
	AllowOrigins []string `yaml:"allowOrigins"`
	IntervalSystemInfoMs int64 `yaml:"intervalSystemInfoMs"`
	SSL          bool
}

/**
Get config from yaml file
*/
func ServerConfigFromFile(filePath string) (*ServerConfig, error) {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	conf := ServerConfig{}
	err = yaml.Unmarshal([]byte(dat), &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
