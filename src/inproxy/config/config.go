package config

import (
	"inproxy/config/parser"
)

type Configurable interface {
	Load(s *parser.Section) error
	Section() string
}

type Config struct {
	HTTP   HTTPConfig
	I2P    I2PConfig
	Filter FilterConfig
}

func (c *Config) Load(fname string) (err error) {
	var conf *parser.Configuration
	conf, err = parser.Read(fname)
	if err == nil {
		configurables := []Configurable{
			&c.HTTP,
			&c.I2P,
			&c.Filter,
		}
		for _, cfg := range configurables {
			s, _ := conf.Section(cfg.Section())
			err = cfg.Load(s)
			if err != nil {
				return
			}
		}
	}
	return
}
