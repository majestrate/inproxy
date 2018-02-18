package config

import (
	"inproxy/config/parser"
)

type HTTPConfig struct {
	BindAddr string
}

func (c *HTTPConfig) Section() string {
	return "http"
}

func (c *HTTPConfig) Load(s *parser.Section) error {
	return nil
}
