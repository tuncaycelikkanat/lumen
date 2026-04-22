package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Rule represents a single detection rule.
type Rule struct {
	ID       int    `yaml:"id"`
	Name     string `yaml:"name"`
	Keyword  string `yaml:"keyword,omitempty"`
	Regex    string `yaml:"regex,omitempty"`
	Severity string `yaml:"severity"`
}

// Config represents the application configuration.
type Config struct {
	Rules []Rule `yaml:"rules"`
}

// LoadConfig reads and parses the YAML configuration file.
func LoadConfig(path string) (*Config, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(buf, &config)
	return &config, err
}
