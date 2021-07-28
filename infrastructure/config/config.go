package config

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

// Config contains common app configuration parameters.
type Config struct {
	Seed   string `yaml:"seed"`
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
}

// ServerAddr compiles the server address.
func (c Config) ServerAddr() string {
	return fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port)
}

// Parse creates configuration from the provided bytes.
func Parse(c io.Reader) (*Config, error) {
	var out Config
	err := yaml.NewDecoder(c).Decode(&out)

	return &out, err
}
