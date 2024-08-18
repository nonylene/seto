package main

import (
	"flag"
	"log"

	"github.com/nonylene/seto/src/seto"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Path to the config file. ${XDG_CONFIG_HOME}/seto/config.json when nothing provided.")
}

func main() {
	flag.Parse()

	cfg, err := seto.ParseConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to parse config: %+v", err)
	}

	log.Fatalf("%+v", seto.Serve(cfg))
}
