package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Address  string `yaml:"address"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
		Timeout  int    `yaml:"timeout"`
	} `yaml:"database"`
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

func NewConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {

		return nil, err
	}
	defer file.Close()
	cfg := &Config{}
	yd := yaml.NewDecoder(file)
	err = yd.Decode(cfg)

	if err != nil {
		return nil, err
	}
	return cfg, nil
}
