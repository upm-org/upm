package config

/*
	UPM Config loader

	Loads configuration files from listed sources
*/

import (
	"upm/logger"

	"gopkg.in/gcfg.v1"
)

type UPMConfig struct {
	BridgesPriorities struct {
		Native int
		Debian int
	}
	logger struct {
		Level int
	}
	Cache struct {
		Dir string
	}
}

var Config UPMConfig

func ReadConfig(fileURL string) error {
	logger.SetLevel(logger.DEBUG)
	logger.SetPrefix("config: ")

	logger.Debugf("Parsing config file %s", fileURL)

	if err := gcfg.ReadFileInto(&Config, fileURL); err != nil {
		return err
	}

	logger.Debugf("Successfully parsed %s", fileURL)

	return nil
}

