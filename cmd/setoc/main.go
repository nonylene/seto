package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nonylene/seto/src/setoc"
)

func die() {
	log.Fatalf("Subcommand \"browser\" or \"code\" is required")
}

func browser(args []string) error {
	fs := flag.NewFlagSet("browser", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage: browser [flags] {url}\n")
		fs.PrintDefaults()
	}

	cfgPath := fs.String("config", "", "Path to the config file. ${XDG_CONFIG_HOME}/seto/config.json when nothing provided")
	err := fs.Parse(args)
	if err != nil {
		// Unexpected
		return err
	}

	if fs.NArg() != 1 {
		return errors.New("a url is required as argument")
	}
	url := fs.Arg(0)

	cfg, err := setoc.ParseConfig(*cfgPath)
	if err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	return setoc.Browser(cfg, url)
}

func code(args []string) error {
	fs := flag.NewFlagSet("code", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage: code [flags] {path}\n")
		fs.PrintDefaults()
	}

	cfgPath := fs.String("config", "", "Path to the config file. ${XDG_CONFIG_HOME}/seto/config.json when nothing provided")
	devContainer := fs.Bool("d", false, "Open in devContainer")
	err := fs.Parse(args)
	if err != nil {
		// Unexpected
		return err
	}

	if fs.NArg() != 1 {
		return errors.New("a path is required as argument")
	}
	path := fs.Arg(0)

	cfg, err := setoc.ParseConfig(*cfgPath)
	if err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	return setoc.Code(cfg, path, *devContainer)
}

func main() {
	if len(os.Args) < 2 {
		die()
	}
	flag.Usage = func() {
		fmt.Print("aa")
		flag.PrintDefaults()
	}

	subArgs := os.Args[2:]
	var err error
	switch os.Args[1] {
	case "browser":
		err = browser(subArgs)
	case "code":
		err = code(subArgs)
	default:
		die()
	}

	if err != nil {
		log.Fatal(err)
	}
}
