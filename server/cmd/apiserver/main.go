package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Gugush284/Go-server.git/internal/apiserver/apiserver"
)

var configPath string
var sessionKeyPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	flag.StringVar(&sessionKeyPath, "sessionKey-Path", "configs/SessionKey.toml", "path to SessionKey")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	_, err = toml.DecodeFile(sessionKeyPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
