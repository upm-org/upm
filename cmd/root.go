package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:	 "upm",
	Short: "UPM is an universal package manager for all Linux distributions",
	Long: `UPM is an universal package manager created
		to reduce amount of effort to maintain packages and to
		be fast and easy for daily use.`,
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
