package config

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	JsonRPCPort string `yaml:"json_rpc_port"`
}

// Config contains common app configuration parameters.
type Config struct {
	Seed   string `yaml:"seed"`
	Server Server `yaml:"server"`
}

// Addr compiles the server address.
func (c Server) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

// RPCAddr compiles the server address for json rpc.
func (c Server) RPCAddr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

// Parse creates configuration from the provided bytes.
func Parse(c io.Reader) (*Config, error) {
	var out Config
	err := yaml.NewDecoder(c).Decode(&out)

	return &out, err
}
