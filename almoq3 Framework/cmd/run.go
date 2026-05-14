package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the almoq3 application server",
	Run: func(cmd *cobra.Command, args []string) {
		if !IsProjectRoot() {
			color.Red("❌ Error: This command can only be executed from the root directory of an almoq3 project.")
			return
		}

		fmt.Printf("%s Starting server...\n", color.CyanString("🚀"))

		// Execute go run main.go
		c := exec.Command("go", "run", "main.go")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		
		if err := c.Run(); err != nil {
			fmt.Printf("%s Server stopped: %v\n", color.RedString("⏹️"), err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
