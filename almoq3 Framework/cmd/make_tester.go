package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var makeTestCmd = &cobra.Command{
	Use:   "make:test [Name]",
	Short: "Create a new automated test",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !IsProjectRoot() {
			color.Red("❌ Error: This command can only be executed from the root directory of an almoq3 project.")
			return
		}

		name := strings.Title(args[0])
		targetPath := fmt.Sprintf("tests/%s_test.go", strings.ToLower(name))
		projectName := getProjectName()

		fmt.Printf("%s Generating Test: %s...\n", color.CyanString("🧪"), color.YellowString(name))

		data := map[string]string{
			"Name":        name,
			"ProjectName": projectName,
		}
		if err := generateFromTemplate("test.tmpl", targetPath, data); err != nil {
			fmt.Printf("%s Failed: %v\n", color.RedString("❌"), err)
			return
		}

		fmt.Printf("%s Test created successfully at %s\n", color.GreenString("✅"), color.CyanString(targetPath))
	},
}

func init() {
	rootCmd.AddCommand(makeTestCmd)
}
