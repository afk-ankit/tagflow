package cmd

import (
	"fmt"
	"log"

	"github.com/ankitsharma/tagflow/internal/db"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		// Initialize DB connection
		db.InitDB()

		project := db.Project{
			Name: projectName,
		}

		result := db.DB.Create(&project)
		if result.Error != nil {
			log.Fatalf("Failed to create project: %v", result.Error)
		}

		fmt.Printf("Project '%s' initialized successfully!\n", projectName)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
