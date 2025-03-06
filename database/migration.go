package database

import (
	"fmt"
	"log"

	"github.com/jamilcse13/go-todo/models"
	"gorm.io/gorm"
)

// AutoMigrateTables runs the migration
func AutoMigrateTables(db *gorm.DB) {
	err := db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatal("❌ Failed to migrate tables:", err)
	}
	fmt.Println("✅ Database migration successful!")
}
