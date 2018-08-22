package main

import (
	"flag"

	"drweb-test/config"
	"drweb-test/google-client"
	"drweb-test/http"
)

var (
	flagConfig = flag.String("config-file", "", "config file name")
)

func main() {
	flag.Parse()

	cfg := new(config.Config)
	config.ReadConfig(*flagConfig, cfg)

	googleClient := google_client.New(cfg.Google)

	server := http.New(cfg.Service.Port, googleClient)
	server.Start()
}
