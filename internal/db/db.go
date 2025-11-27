package db

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Failed to get user home directory:", err)
	}

	dbPath := filepath.Join(homeDir, ".tagflow", "tagflow.db")
	err = os.MkdirAll(filepath.Dir(dbPath), 0755)
	if err != nil {
		log.Fatal("Failed to create database directory:", err)
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	err = DB.AutoMigrate(&Project{}, &Branch{}, &Tag{}, &Deployment{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
