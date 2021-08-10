package config

import (
	"cobragingorm/internal/app/config"

	"github.com/BurntSushi/toml"
)

var defaultConfig = `
debug = true # debug mode flag
runtime = "runtime" # runtime folder

[server]
addr = ":9000" # Listen port
baseUrl = "http://api.example.com" # deployment basic domain

[cors]
origin = ["*"]
methods = ["GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"]
headers = ["Content-Length", "Content-Type", "Origin"]
credentials = true
maxAge = 24 # unit : hours
`

func init() {
	_, err := toml.Decode(defaultConfig, config.Conf)
	if err != nil {
		panic("default config synctax error")
	}
}
