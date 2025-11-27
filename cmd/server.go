package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ankitsharma/tagflow/internal/db"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the Tagflow API server",
	Run: func(cmd *cobra.Command, args []string) {
		db.InitDB()

		http.HandleFunc("/api/projects", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			var projects []db.Project
			db.DB.Preload("Branches").Find(&projects)
			json.NewEncoder(w).Encode(projects)
		})

		http.HandleFunc("/api/branches", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			var branches []db.Branch
			db.DB.Preload("Tags").Find(&branches)
			json.NewEncoder(w).Encode(branches)
		})

		http.HandleFunc("/api/tags", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			var tags []db.Tag
			db.DB.Preload("Branch").Find(&tags)
			json.NewEncoder(w).Encode(tags)
		})

		http.HandleFunc("/api/deployments", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			var deployments []db.Deployment
			db.DB.Preload("Tag").Preload("Tag.Branch").Find(&deployments)
			json.NewEncoder(w).Encode(deployments)
		})

		port := ":8080"
		fmt.Printf("Starting server on %s\n", port)
		log.Fatal(http.ListenAndServe(port, nil))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
