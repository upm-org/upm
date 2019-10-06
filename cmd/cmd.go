package cmd

import (
	"errors"
	"upm/config"
	"upm/logger"

	"github.com/spf13/cobra"
)

var confPath string
var logLevel int

var ErrUnsupportedLog = errors.New("unsupported logging level")

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&confPath, "config", "config/default.conf", "Pass path to UPM config file")
	rootCmd.PersistentFlags().IntVar(&logLevel, "log", 2, "Set log level 0-NONE, 1-INFO, 2-DEBUG")

	rootCmd.AddCommand(unpackCmd)
}

func initConfig() {
	if err := config.ReadConfig(confPath); err != nil {
		logger.Fatal(err)
	}
	if logLevel < logger.NONE || logLevel > logger.DEBUG {
		logger.Fatal(ErrUnsupportedLog)
	} else {
		logger.SetLevel(logLevel)
	}
}

func Execute() error {
	//rootCmd.RunE(rootCmd, os.Args)
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
