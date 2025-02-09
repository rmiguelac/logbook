package main

import (
	"log"

	"github.com/rmiguelac/logbook/backend/internal/database"
	"github.com/rmiguelac/logbook/backend/internal/models"
	"github.com/rmiguelac/logbook/backend/pkg/config"
)

func main() {
	cfg := config.Load()
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}

	// Explicit migration call using db
	err = db.AutoMigrate(
		&models.Task{},
		&models.Comment{},
		&models.TaskHistory{},
		&models.Note{},
		&models.Tag{},
		&models.User{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Database migrated successfully!")
}
