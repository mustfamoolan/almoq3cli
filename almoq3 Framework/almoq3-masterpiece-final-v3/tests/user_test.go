package tests

import (
	"testing"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"mytest/app/models"
)

func SetupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	return db
}

func TestUser(t *testing.T) {
	db := SetupTestDB()
	
	t.Run("Check if test DB is ready", func(t *testing.T) {
		if db == nil {
			t.Errorf("Database connection failed")
		}
	})
}
