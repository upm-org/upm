package main

import (
	"upm/cmd"
	"upm/logger"
)

func main() {
	logger.SetLevel(logger.INFO)
	logger.SetPrefix("main: ")

	if err := cmd.Execute(); err != nil {
		logger.Fatal(err)
	}
}
