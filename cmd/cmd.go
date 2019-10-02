package cmd

import "upm/logger"

func init() {
	rootCmd.PersistentFlags().IntVar(&logger.Log.Level, "log", 1, "Set log level 0-NONE, 1-INFO, 2-DEBUG")

	rootCmd.AddCommand(unpackCmd)
}

