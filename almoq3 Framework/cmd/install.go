package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install almoq3 CLI globally on your system",
	Run: func(cmd *cobra.Command, args []string) {
		exePath, err := os.Executable()
		if err != nil {
			fmt.Printf("%s %v\n", color.RedString("Error:"), err)
			return
		}

		dir := filepath.Dir(exePath)
		fmt.Printf("%s Installing almoq3 CLI from: %s\n", color.CyanString("⚙️"), dir)

		// PowerShell command to add to User PATH permanently
		psCommand := fmt.Sprintf(`
			$currentPath = [Environment]::GetEnvironmentVariable("Path", "User");
			if ($currentPath -notlike "*%s*") {
				[Environment]::SetEnvironmentVariable("Path", $currentPath + ";%s", "User");
				Write-Host "PATH updated successfully.";
			} else {
				Write-Host "PATH already contains this directory.";
			}
		`, strings.ReplaceAll(dir, `\`, `\\`), strings.ReplaceAll(dir, `\`, `\\`))

		out, err := exec.Command("powershell", "-Command", psCommand).CombinedOutput()
		if err != nil {
			fmt.Printf("%s Failed to update PATH: %v\n", color.RedString("❌"), err)
			fmt.Println(string(out))
			return
		}

		fmt.Printf("%s %s", color.GreenString("✅"), string(out))
		fmt.Printf("\n%s Please %s your Terminal/CMD for changes to take effect.\n", 
			color.YellowString("IMPORTANT:"), color.HiWhiteString("RESTART"))
		fmt.Printf("After restart, you can use %s from any folder!\n", color.CyanString("almoq3"))
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
