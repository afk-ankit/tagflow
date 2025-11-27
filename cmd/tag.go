package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/ankitsharma/tagflow/internal/db"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Manage tags",
}

var listTagCmd = &cobra.Command{
	Use:   "list",
	Short: "List tags",
	Run: func(cmd *cobra.Command, args []string) {
		db.InitDB()
		var tags []db.Tag
		db.DB.Preload("Branch").Find(&tags)

		fmt.Println("Tags:")
		for _, tag := range tags {
			fmt.Printf("Name: %s, Branch: %s, CreatedBy: %s\n", tag.Name, tag.Branch.Name, tag.CreatedBy)
		}
	},
}

var suggestTagCmd = &cobra.Command{
	Use:   "suggest",
	Short: "Suggest next tag",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic to suggest tag based on previous tags
		// For simplicity, let's just find the latest tag and increment
		db.InitDB()
		var lastTag db.Tag
		db.DB.Order("created_at desc").First(&lastTag)

		if lastTag.Name == "" {
			fmt.Println("Suggested Tag: v0.0.1")
			return
		}

		// Simple increment logic (this is a placeholder, real logic would be more complex)
		fmt.Printf("Last Tag: %s. Suggested Tag: %s-next\n", lastTag.Name, lastTag.Name)
	},
}

var createTagCmd = &cobra.Command{
	Use:   "create [tag-name]",
	Short: "Create a new tag",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tagName := args[0]

		// Get current branch
		out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
		if err != nil {
			log.Fatalf("Failed to get current branch: %v", err)
		}
		currentBranch := strings.TrimSpace(string(out))

		db.InitDB()
		var branch db.Branch
		if err := db.DB.Where("name = ?", currentBranch).First(&branch).Error; err != nil {
			log.Printf("Warning: Current branch '%s' is not tracked in Tagflow DB.", currentBranch)
		}

		// Create git tag
		gitCmd := exec.Command("git", "tag", tagName)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		if err := gitCmd.Run(); err != nil {
			log.Fatalf("Failed to create git tag: %v", err)
		}

		// Save to DB
		tag := db.Tag{
			Name:      tagName,
			BranchID:  branch.ID,
			CreatedBy: "user", // TODO: Get actual user
		}
		if err := db.DB.Create(&tag).Error; err != nil {
			log.Fatalf("Failed to track tag in DB: %v", err)
		}

		fmt.Printf("Tag '%s' created and tracked.\n", tagName)
	},
}

func init() {
	tagCmd.AddCommand(listTagCmd)
	tagCmd.AddCommand(suggestTagCmd)
	tagCmd.AddCommand(createTagCmd)
	rootCmd.AddCommand(tagCmd)
}
