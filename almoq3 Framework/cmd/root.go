package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/almoq3/almoq3-cli/internal/templates"
	"github.com/spf13/cobra"
)

const Version = "1.0.1"

var rootCmd = &cobra.Command{
	Use:     "almoq3",
	Version: Version,
	Short:   "almoq3 is a developer-friendly Go framework CLI",
	Long:    `A high-performance CLI tool for scaffolding and managing almoq3 Framework projects.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func IsProjectRoot() bool {
	cwd, _ := os.Getwd()
	envPath := filepath.Join(cwd, ".env")
	modPath := filepath.Join(cwd, "go.mod")

	_, errEnv := os.Stat(envPath)
	_, errMod := os.Stat(modPath)

	return errEnv == nil && errMod == nil
}

func init() {
	// Root flags if needed
}

func generateFromTemplate(tmplName, targetPath string, data interface{}) error {
	content, err := templates.Files.ReadFile(tmplName)
	if err != nil {
		return err
	}

	tmpl, err := template.New(tmplName).Parse(string(content))
	if err != nil {
		return err
	}

	os.MkdirAll(filepath.Dir(targetPath), 0755)

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	return os.WriteFile(targetPath, buf.Bytes(), 0644)
}

func getProjectName() string {
	// In a real implementation, we would parse go.mod
	return "mytest"
}
