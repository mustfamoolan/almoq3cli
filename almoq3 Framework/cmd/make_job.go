package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var makeJobCmd = &cobra.Command{
	Use:   "make:job [Name]",
	Short: "Create a new background job",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !IsProjectRoot() {
			color.Red("❌ Error: This command can only be executed from the root directory of an almoq3 project.")
			return
		}

		name := strings.Title(args[0])
		targetPath := fmt.Sprintf("app/jobs/%s_job.go", strings.ToLower(name))

		fmt.Printf("%s Generating Job: %s...\n", color.CyanString("⚡"), color.YellowString(name))

		data := map[string]string{"Name": name}
		if err := generateFromTemplate("job.tmpl", targetPath, data); err != nil {
			fmt.Printf("%s Failed: %v\n", color.RedString("❌"), err)
			return
		}

		fmt.Printf("%s Job created successfully at %s\n", color.GreenString("✅"), color.CyanString(targetPath))
	},
}

func init() {
	rootCmd.AddCommand(makeJobCmd)
}
