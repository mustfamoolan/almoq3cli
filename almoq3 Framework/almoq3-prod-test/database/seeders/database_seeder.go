package seeders

import (
	"fmt"
	"github.com/fatih/color"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// Register your seeders here
		// Example:
		// if err := runSeeder(tx, &UserSeeder{}); err != nil { return err }
		
		fmt.Printf("%s All seeders completed successfully.\n", color.GreenString("🏁"))
		return nil
	})
}

func runSeeder(db *gorm.DB, seeder interface{ Run(db *gorm.DB) error }) error {
	name := fmt.Sprintf("%T", seeder)
	fmt.Printf("%s Running Seeder: %s...\n", color.CyanString("⚡"), color.YellowString(name))
	
	if err := seeder.Run(db); err != nil {
		fmt.Printf("%s Seeder %s failed: %v\n", color.RedString("❌"), name, err)
		return err
	}
	
	fmt.Printf("%s Seeder %s completed.\n", color.GreenString("✅"), name)
	return nil
}
