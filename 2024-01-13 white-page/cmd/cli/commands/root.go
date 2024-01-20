package commands

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "the-write-page-cli",
		Short: "The write page applications",
		Long:  ``,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// do something
}

func Execute() error {
	return rootCmd.Execute()
}
