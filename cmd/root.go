package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tagflow",
	Short: "Tagflow is a CLI tool for managing git tags and deployments",
	Long: `Tagflow helps you manage your git tags, feature branches, and track deployments 
across different environments. It integrates with your existing git workflow.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Tagflow CLI v0.1.0")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
