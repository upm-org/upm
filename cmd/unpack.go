package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"upm/pkg/manager"
)

var unpackCmd = &cobra.Command{
	Use: "unpack",
	Short: "Unpack packages",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("too few arguments, need at least source and destination")
		}

		to := args[len(args) - 1]
		sources := args[:len(args)-1]
		errChan := make(chan error, len(sources))

		for _, file := range sources {
			go func(filePath string) {
				errChan <- manager.Unpack(filePath, to)
			}(file)
		}
		for i := 0; i < len(sources); i++ {
			if err := <-errChan; err != nil {
				return err
			}
		}
		close(errChan)

		return nil
	},
}
