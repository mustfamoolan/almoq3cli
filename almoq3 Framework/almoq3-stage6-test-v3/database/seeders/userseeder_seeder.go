package seeders

import (
	"gorm.io/gorm"
)

type UserseederSeeder struct{}

func (s *UserseederSeeder) Run(db *gorm.DB) error {
	// Write your seeding logic here
	// Example:
	// db.Create(&models.User{Name: "John Doe"})

	return nil
}
