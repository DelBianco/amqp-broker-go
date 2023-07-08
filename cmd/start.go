package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start Broker",
		Long:  `Start Broker`,
		Run:   run,
	}
)

func init() {
	fmt.Println("Start initialization")
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println("Started")
}
