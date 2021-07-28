package main

import (
	"flag"
	"os"

	"github.com/trustwallet/wallet-core"

	"github.com/default23/crypto-api/infrastructure/config"
	"github.com/default23/crypto-api/infrastructure/logging"
)

func main() {
	var configPath string
	logger := logging.New()

	flag.StringVar(&configPath, "config", "", "path to configuration file")
	flag.Parse()

	if configPath == "" {
		panic("config path is not specified. Use '--config=/path/to/config' parameter to run application")
	}

	configFile, err := os.OpenFile(configPath, os.O_RDONLY, 0)
	if err != nil {
		logger.Fatal(err)
		return
	}

	cfg, err := config.Parse(configFile)
	if err != nil {
		logger.Fatal(err)
		return
	}

	logger.Info(cfg)
}
