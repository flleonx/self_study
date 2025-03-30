package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "worker",
		Short: "CLI to run workers",
		Long:  ``,
	}

	rootCmd.AddCommand(ExecutePocWorker())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
