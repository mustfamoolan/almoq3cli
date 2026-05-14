package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/almoq3/almoq3-cli/internal/scaffolder"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [projectName]",
	Short: "Create a new almoq3 project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		// Validation: No special characters
		if !isValidProjectName(projectName) {
			color.Red("Error: Project name can only contain letters, numbers, hyphens, and underscores.")
			return
		}

		err := scaffolder.Scaffold(projectName)
		if err != nil {
			fmt.Printf("%s %v\n", color.RedString("Error:"), err)
			return
		}

		// 2. Run go mod tidy automatically
		fmt.Printf("%s Tidying dependencies...\n", color.CyanString("⚙️"))
		tidyCmd := exec.Command("go", "mod", "tidy")
		tidyCmd.Dir = projectName
		if err := tidyCmd.Run(); err != nil {
			fmt.Printf("%s Warning: Failed to tidy dependencies: %v\n", color.YellowString("⚠️"), err)
		}

		// 3. Auto-Install Check: Try to make the CLI global if not already
		exePath, _ := os.Executable()
		dir := filepath.Dir(exePath)
		psCommand := fmt.Sprintf(`
			$currentPath = [Environment]::GetEnvironmentVariable("Path", "User");
			if ($currentPath -notlike "*%s*") {
				[Environment]::SetEnvironmentVariable("Path", $currentPath + ";%s", "User");
			}
		`, strings.ReplaceAll(dir, `\`, `\\`), strings.ReplaceAll(dir, `\`, `\\`))
		_ = exec.Command("powershell", "-Command", psCommand).Run()
	},
}

func isValidProjectName(name string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
	return re.MatchString(name)
}

func init() {
	rootCmd.AddCommand(newCmd)
}
