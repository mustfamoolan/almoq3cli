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

var makeServiceCmd = &cobra.Command{
	Use:   "make:service [Name]",
	Short: "Create a new business service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !IsProjectRoot() {
			color.Red("❌ Error: This command can only be executed from the root directory of an almoq3 project.")
			return
		}

		name := args[0]
		// Capitalize first letter
		structName := strings.Title(strings.ToLower(name))
		if !strings.HasSuffix(structName, "Service") {
			structName += "Service"
		}
		
		fileName := strings.ToLower(name) + "_service.go"
		targetPath := filepath.Join("app", "services", fileName)

		fmt.Printf("%s Generating Service: %s...\n", color.CyanString("⚡"), color.YellowString(structName))

		data, err := templates.Files.ReadFile("service.tmpl")
		if err != nil {
			fmt.Printf("%s %v\n", color.RedString("Error:"), err)
			return
		}

		tmpl, err := template.New("service").Parse(string(data))
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

		fmt.Printf("%s Service created successfully at %s\n", color.GreenString("✅"), color.CyanString(targetPath))
	},
}

func init() {
	rootCmd.AddCommand(makeServiceCmd)
}
