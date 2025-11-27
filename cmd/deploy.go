package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/ankitsharma/tagflow/internal/db"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy [tag-name] [environment]",
	Short: "Deploy a tag to an environment",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		tagName := args[0]
		env := args[1]

		db.InitDB()
		var tag db.Tag
		if err := db.DB.Where("name = ?", tagName).First(&tag).Error; err != nil {
			log.Fatalf("Tag '%s' not found.", tagName)
		}

		// In a real scenario, this would trigger a deployment pipeline or similar.
		// For now, we just track it.
		fmt.Printf("Deploying tag '%s' to '%s'...\n", tagName, env)

		deployment := db.Deployment{
			TagID:       tag.ID,
			Environment: env,
			DeployedBy:  "user", // TODO: Get actual user
			DeployedAt:  time.Now(),
		}

		if err := db.DB.Create(&deployment).Error; err != nil {
			log.Fatalf("Failed to track deployment: %v", err)
		}

		fmt.Println("Deployment tracked successfully.")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
