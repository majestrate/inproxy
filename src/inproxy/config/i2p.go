package config

import (
	"inproxy/config/parser"
)

type I2PConfig struct {
}

func (c *I2PConfig) Section() string {
	return "i2p"
}

func (c *I2PConfig) Load(s *parser.Section) error {
	return nil
}
