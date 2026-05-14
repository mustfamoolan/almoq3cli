package migrations

import (
	"gorm.io/gorm"
	"time"
)

type CreateUsersTable struct{}

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"size:255"`
	Email     string         `gorm:"uniqueIndex;size:255"`
	Password  string         `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *CreateUsersTable) Up(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

func (m *CreateUsersTable) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&User{})
}
