package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "simple-message-broker",
		Short: "Go implementation of a simple-message-broker broker",
		Long:  `simple-message-broker is a lightweight broker with limited features`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(startCmd)
}
