package database

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Migration struct {
	ID        uint   `gorm:"primaryKey"`
	Migration string `gorm:"unique;not null"`
	Batch     int    `gorm:"not null"`
	CreatedAt time.Time
}

func Connect(driver, host, port, database, username, password string) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch driver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username, password, host, port, database)
		dialector = mysql.Open(dsn)
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host, username, password, database, port)
		dialector = postgres.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(database + ".db")
	default:
		return nil, fmt.Errorf("unsupported driver: %s", driver)
	}

	return gorm.Open(dialector, &gorm.Config{})
}

func Migrate(db *gorm.DB, migrationsPath string) error {
	// Ensure migrations table exists
	if err := db.AutoMigrate(&Migration{}); err != nil {
		return err
	}

	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		return err
	}

	var pendingFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".up.sql") {
			var count int64
			db.Model(&Migration{}).Where("migration = ?", file.Name()).Count(&count)
			if count == 0 {
				pendingFiles = append(pendingFiles, file.Name())
			}
		}
	}

	if len(pendingFiles) == 0 {
		fmt.Printf("%s All migrations are up to date. Nothing to run.\n", color.YellowString("✨"))
		return nil
	}

	sort.Strings(pendingFiles)

	var lastBatch int
	db.Model(&Migration{}).Select("max(batch)").Scan(&lastBatch)
	currentBatch := lastBatch + 1

	fmt.Printf("%s Found %d pending migrations. Starting execution...\n", color.CyanString("📁"), len(pendingFiles))

	for _, fileName := range pendingFiles {
		fmt.Printf("%s Executing: %s\n", color.YellowString("⚡"), fileName)
		content, err := os.ReadFile(filepath.Join(migrationsPath, fileName))
		if err != nil {
			return err
		}

		err = db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Exec(string(content)).Error; err != nil {
				return err
			}
			return tx.Create(&Migration{Migration: fileName, Batch: currentBatch}).Error
		})

		if err != nil {
			fmt.Printf("%s Migration failed: %s. Rolling back transaction...\n", color.RedString("❌"), fileName)
			return fmt.Errorf("failed to execute %s: %w", fileName, err)
		}
		fmt.Printf("  %s Success: %s\n", color.GreenString("✅"), fileName)
	}

	fmt.Printf("%s Migrations completed successfully.\n", color.GreenString("🏁"))
	return nil
}

func Rollback(db *gorm.DB, migrationsPath string) error {
	var lastBatch int
	db.Model(&Migration{}).Select("max(batch)").Scan(&lastBatch)

	if lastBatch == 0 {
		fmt.Printf("%s No migration history found. Nothing to rollback.\n", color.YellowString("✨"))
		return nil
	}

	var migrations []Migration
	db.Where("batch = ?", lastBatch).Order("id desc").Find(&migrations)

	fmt.Printf("%s Rolling back %d migrations from batch %d...\n", color.YellowString("⏪"), len(migrations), lastBatch)

	for _, m := range migrations {
		downFile := strings.Replace(m.Migration, ".up.sql", ".down.sql", 1)
		fmt.Printf("%s Rolling back: %s\n", color.YellowString("⚡"), downFile)

		content, err := os.ReadFile(filepath.Join(migrationsPath, downFile))
		if err != nil {
			return err
		}

		err = db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Exec(string(content)).Error; err != nil {
				return err
			}
			return tx.Delete(&m).Error
		})

		if err != nil {
			fmt.Printf("%s Rollback failed: %s. Rolling back transaction...\n", color.RedString("❌"), downFile)
			return fmt.Errorf("failed to rollback %s: %w", downFile, err)
		}
		fmt.Printf("  %s Success: %s\n", color.GreenString("✅"), downFile)
	}

	fmt.Printf("%s Rollback completed successfully.\n", color.GreenString("🏁"))
	return nil
}
