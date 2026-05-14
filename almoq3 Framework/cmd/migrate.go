package cmd

import (
	"fmt"

	"github.com/almoq3/almoq3-cli/internal/database"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run pending migrations",
	Run: func(cmd *cobra.Command, args []string) {
		if !IsProjectRoot() {
			color.Red("❌ Error: This command can only be executed from the root directory of an almoq3 project.")
			return
		}

		db, err := connectToDB()
		if err != nil {
			fmt.Printf("%s %v\n", color.RedString("Error:"), err)
			return
		}

		err = database.Migrate(db, "database/migrations")
		if err != nil {
			fmt.Printf("%s %v\n", color.RedString("Error:"), err)
			return
		}
	},
}

var rollbackCmd = &cobra.Command{
	Use:   "migrate:rollback",
	Short: "Rollback the last batch of migrations",
	Run: func(cmd *cobra.Command, args []string) {
		if !IsProjectRoot() {
			color.Red("❌ Error: This command can only be executed from the root directory of an almoq3 project.")
			return
		}

		db, err := connectToDB()
		if err != nil {
			fmt.Printf("%s %v\n", color.RedString("Error:"), err)
			return
		}

		err = database.Rollback(db, "database/migrations")
		if err != nil {
			fmt.Printf("%s %v\n", color.RedString("Error:"), err)
			return
		}
	},
}

func connectToDB() (*gorm.DB, error) {
	// Load .env from current directory
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read .env file: %w", err)
	}

	driver := viper.GetString("DB_DRIVER")
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	dbName := viper.GetString("DB_DATABASE")
	user := viper.GetString("DB_USERNAME")
	pass := viper.GetString("DB_PASSWORD")

	if driver == "" {
		return nil, fmt.Errorf("DB_DRIVER is missing in .env")
	}

	return database.Connect(driver, host, port, dbName, user, pass)
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(rollbackCmd)
}
