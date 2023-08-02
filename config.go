package main

import (
	"github.com/jinzhu/configor"
)

type Configuration struct {
	Endpoint string `required:"true" env:"ENDPOINT"`
	Username string `default:"incoming-webhook" env:"USERNAME"`
	Icon     string `default:":ghost:" env:"ICON"`
}

var Config = new(Configuration)

func init() {
	configor.Load(Config)
}
