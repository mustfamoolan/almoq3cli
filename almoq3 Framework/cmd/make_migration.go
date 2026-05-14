package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var makeMigrationCmd = &cobra.Command{
	Use:   "make:migration [name]",
	Short: "Create a new migration file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !IsProjectRoot() {
			color.Red("❌ Error: This command can only be executed from the root directory of an almoq3 project.")
			return
		}

		name := args[0]
		timestamp := time.Now().Format("20060102150405")
		dir := "database/migrations"

		// Ensure directory exists
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("%s %v\n", color.RedString("Error:"), err)
			return
		}

		upFile := filepath.Join(dir, fmt.Sprintf("%s_%s.up.sql", timestamp, name))
		downFile := filepath.Join(dir, fmt.Sprintf("%s_%s.down.sql", timestamp, name))

		// Create up file
		if err := os.WriteFile(upFile, []byte("-- Up migration"), 0644); err != nil {
			fmt.Printf("%s %v\n", color.RedString("Error:"), err)
			return
		}

		// Create down file
		if err := os.WriteFile(downFile, []byte("-- Down migration"), 0644); err != nil {
			fmt.Printf("%s %v\n", color.RedString("Error:"), err)
			return
		}

		fmt.Printf("%s Created migration files:\n", color.GreenString("✓"))
		fmt.Printf("  - %s\n", upFile)
		fmt.Printf("  - %s\n", downFile)
	},
}

func init() {
	rootCmd.AddCommand(makeMigrationCmd)
}
