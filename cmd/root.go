package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"upm/config"
	"upm/logger"
)

var rootCmd = &cobra.Command{
	Use: "upm",
	Short: "UPM is an universal package manager for all Linux distributions",
	Long: `UPM is an universal package manager created
to reduce amount of effort to maintain packages and to
be fast and easy for daily use.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.SetLevel(logLevel)

		if err := config.ReadConfig(confPath); err != nil {
			return err
		}

		// Creating cache directory if doesn't exist
		if err := os.MkdirAll(config.Config.Cache.Dir, 0755); err != nil {
			return err
		}

		return nil
	},
}
