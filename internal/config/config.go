package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Rule struct {
	ID       int    `yaml:"id"`
	Name     string `yaml:"name"`
	Keyword  string `yaml:"Keyword"`
	Severity string `yaml:"severity"`
}
type Config struct {
	Rules []Rule `yaml:"rules"`
}

func LoadConfig(path string) (*Config, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(buf, &config)
	return &config, err
}
