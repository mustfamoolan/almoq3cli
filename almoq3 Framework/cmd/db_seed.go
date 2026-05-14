package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var dbSeedCmd = &cobra.Command{
	Use:   "db:seed",
	Short: "Run the database seeders",
	Run: func(cmd *cobra.Command, args []string) {
		if !IsProjectRoot() {
			color.Red("❌ Error: This command can only be executed from the root directory of an almoq3 project.")
			return
		}

		fmt.Printf("%s Starting Database Seeding...\n", color.CyanString("⚙️"))

		// Run the seeder runner via go run
		// This ensures we run the code within the user's project context
		process := exec.Command("go", "run", "cmd/seed/main.go")
		process.Stdout = os.Stdout
		process.Stderr = os.Stderr

		if err := process.Run(); err != nil {
			fmt.Printf("%s Seeding process failed: %v\n", color.RedString("❌"), err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(dbSeedCmd)
}
