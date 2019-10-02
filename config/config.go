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
	Log struct {
		Level int
	}
	Cache struct {
		Dir string
	}
}

var Config UPMConfig

func (conf *UPMConfig) ReadConfig(fileURL string) {
	Log := logger.Log
	Log.Level = logger.INFO
	Log.Debug("Parsing config file %s\n", fileURL)
	if err := gcfg.ReadFileInto(conf, fileURL); err != nil {
		Log.Fatal("%s", err)
	}
	Log.Debug("Successfully parsed %s\n", fileURL)
}

