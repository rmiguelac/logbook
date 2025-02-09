package database

import (
	"github.com/rmiguelac/logbook/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate models
	err = db.AutoMigrate(
		&models.Task{},
		&models.Comment{},
		&models.TaskHistory{},
		&models.Note{},
		&models.Tag{},
		&models.User{},
	)

	return db, err
}
