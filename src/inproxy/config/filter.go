package config

import (
	"inproxy/config/parser"
)

type FilterConfig struct {
}

func (c *FilterConfig) Section() string {
	return "filter"
}

func (c *FilterConfig) Load(s *parser.Section) error {
	return nil
}
