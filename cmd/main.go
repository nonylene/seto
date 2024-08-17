package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nonylene/seto/src"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Path to the config file. ${XDG_CONFIG_HOME}/seto/config.json when nothing provided.")
}

func main() {
	flag.Parse()

	cfg, err := src.ParseConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to parse config: %+v", err)
	}

	fmt.Printf("%+v", cfg)
}
