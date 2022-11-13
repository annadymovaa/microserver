package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/annadymovaa/avito-test/tree/main/http-rest-api/cmd/inetrnal/app/apiserver"
	//"honnef.co/go/tools/config"
	//"github.com/annadymovaa/avito-test/tree/start/http-rest-api/cmd/inetrnal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
