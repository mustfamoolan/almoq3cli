package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/almoq3/almoq3-cli/internal/templates"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var makeSeederCmd = &cobra.Command{
	Use:   "make:seeder [Name]",
	Short: "Create a new database seeder",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !IsProjectRoot() {
			color.Red("❌ Error: This command can only be executed from the root directory of an almoq3 project.")
			return
		}

		name := args[0]
		// Ensure suffix
		structName := strings.Title(strings.ToLower(name))
		if !strings.HasSuffix(structName, "Seeder") {
			structName += "Seeder"
		}
		
		fileName := strings.ToLower(name) + "_seeder.go"
		targetPath := filepath.Join("database", "seeders", fileName)

		fmt.Printf("%s Generating Seeder: %s...\n", color.CyanString("⚡"), color.YellowString(structName))

		data, err := templates.Files.ReadFile("seeder.tmpl")
		if err != nil {
			fmt.Printf("%s %v\n", color.RedString("Error:"), err)
			return
		}

		tmpl, err := template.New("seeder").Parse(string(data))
		if err != nil {
			fmt.Printf("%s %v\n", color.RedString("Error:"), err)
			return
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, map[string]string{"Name": structName}); err != nil {
			fmt.Printf("%s %v\n", color.RedString("Error:"), err)
			return
		}

		if err := os.WriteFile(targetPath, buf.Bytes(), 0644); err != nil {
			fmt.Printf("%s %v\n", color.RedString("Error:"), err)
			return
		}

		fmt.Printf("%s Seeder created successfully at %s\n", color.GreenString("✅"), color.CyanString(targetPath))
	},
}

func init() {
	rootCmd.AddCommand(makeSeederCmd)
}
