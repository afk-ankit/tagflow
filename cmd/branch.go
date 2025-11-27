package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/ankitsharma/tagflow/internal/db"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Manage branches",
}

var createBranchCmd = &cobra.Command{
	Use:   "create [branch-name]",
	Short: "Create a new feature branch",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branchName := args[0]
		projectName, _ := cmd.Flags().GetString("project")

		db.InitDB()

		var project db.Project
		if err := db.DB.Where("name = ?", projectName).First(&project).Error; err != nil {
			log.Fatalf("Project '%s' not found. Please init project first.", projectName)
		}

		// Create git branch
		// Assuming we are in the repo
		gitCmd := exec.Command("git", "checkout", "-b", branchName)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		if err := gitCmd.Run(); err != nil {
			log.Fatalf("Failed to create git branch: %v", err)
		}

		// Save to DB
		branch := db.Branch{
			Name:      branchName,
			ProjectID: project.ID,
			Type:      "feature",
		}

		if err := db.DB.Create(&branch).Error; err != nil {
			log.Fatalf("Failed to track branch in DB: %v", err)
		}

		fmt.Printf("Branch '%s' created and tracked for project '%s'\n", branchName, projectName)
	},
}

func init() {
	branchCmd.AddCommand(createBranchCmd)
	createBranchCmd.Flags().StringP("project", "p", "", "Project name (required)")
	createBranchCmd.MarkFlagRequired("project")
	rootCmd.AddCommand(branchCmd)
}
