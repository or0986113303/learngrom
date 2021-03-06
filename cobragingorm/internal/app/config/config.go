package config

import (
	"github.com/go-eyas/toolkit/log"
)

type Config struct {
	// running mode
	Debug bool `env:"API_DEBUG"`
	// runtime folder
	Runtime string `default:"runtime"`

	// http service parameters
	Server struct {
		Addr    string `env:"API_ADDR"`
		BaseURL string
	}

	// cors config
	Cors struct {
		Origin      []string
		Methods     []string
		Headers     []string
		Credentials bool
		MaxAge      int
	}

	// log config
	Log log.LogConfig
}

// Conf of global variable
var Conf = &Config{}
