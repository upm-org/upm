package cmd

var confPath string
var logLevel int

func init() {
	rootCmd.PersistentFlags().IntVar(&logLevel, "log", 2, "Set log level 0-NONE, 1-INFO, 2-DEBUG")
	rootCmd.PersistentFlags().StringVar(&confPath, "config", "config/default.conf", "Pass path to UPM config file")

	rootCmd.AddCommand(unpackCmd)
}

func Execute() error {
	//rootCmd.RunE(rootCmd, os.Args)
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
