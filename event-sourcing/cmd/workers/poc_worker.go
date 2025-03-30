package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ExecutePocWorker() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "poc",
		Short: "CLI to run workers",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("POC worker execution")
		},
	}

	return cmd
}
