package main

import (
	"fmt"
	"flag"
	"upm/logger"
	"upm/config"
)

var log logger.UPMLogger
var cfg config.UPMConfig

func init(){
	/*
		Init function

		Parses CLI arguments and init config file
	*/

	var logLvl int
	var configPath string

	flag.IntVar(&logLvl, "log", 1, "Verbose level: 0-NONE, 1-INFO, 2-DEBUG")
	flag.StringVar(&configPath, "config", "./config/default.conf", "Provide custom path to config file")

	flag.Parse()

	log := logger.UPMLogger{
		Lvl: logLvl,
	}
	// Logger injection
	config.Log = log

	var err error
	cfg, err = config.LoadConfig(configPath)

	if err != nil {
		log.Error("Failed to load config: %s", err)
	}
}

func main() {
	fmt.Println(cfg)
}
