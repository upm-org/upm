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
}

var Log logger.UPMLogger

func LoadConfig(fileURL string) (UPMConfig, error) {
	Log.Debug("Parsing config file %s", fileURL)
	var res UPMConfig
	if err := gcfg.ReadFileInto(&res, fileURL); err != nil {
		return UPMConfig{}, err
	}
	Log.Debug("Successfully parsed %s", fileURL)
	return res, nil
}

