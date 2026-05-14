package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/almoq3/almoq3-cli/internal/templates"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const Version = "2.0.3"
const RepoURL = "https://api.github.com/repos/mustfamoolan/almoq3cli/releases/latest"

var rootCmd = &cobra.Command{
	Use:     "almoq3",
	Version: Version,
	Short:   "almoq3 is a developer-friendly Go framework CLI",
	Long:    `A high-performance CLI tool for scaffolding and managing almoq3 Framework projects.`,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// Don't check for updates during the self-update command itself
		if cmd.Name() != "self-update" {
			checkForUpdates()
		}
	},
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

func checkForUpdates() {
	// This is a simple version check. In a high-traffic app, 
	// we would cache this result to avoid hitting GitHub API limits.
	resp, err := http.Get(RepoURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return
	}

	latestVersion := strings.TrimPrefix(release.TagName, "v")
	if latestVersion != "" && latestVersion != Version {
		fmt.Println()
		color.Yellow("┌──────────────────────────────────────────────────────────┐")
		fmt.Printf("%s  Update available: %s → %s %s\n", 
			color.YellowString("│"), 
			color.RedString(Version), 
			color.GreenString(latestVersion),
			color.YellowString("│"))
		fmt.Printf("%s  Run %s to update    %s\n", 
			color.YellowString("│"), 
			color.CyanString("almoq3 self-update"),
			color.YellowString("│"))
		color.Yellow("└──────────────────────────────────────────────────────────┘")
		fmt.Println()
	}
}
