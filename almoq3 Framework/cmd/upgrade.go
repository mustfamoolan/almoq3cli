package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"text/template"

	"github.com/almoq3/almoq3-cli/internal/templates"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// ProjectMeta represents the almoq3.json file structure
type ProjectMeta struct {
	ProjectName      string `json:"project_name"`
	FrameworkVersion string `json:"framework_version"`
	CreatedAt        string `json:"created_at"`
}

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Safely upgrade framework files without touching your custom code",
	Long: `Upgrades only framework-managed files to the latest version.

Protected (NEVER touched):
  app/controllers/, app/models/, app/services/
  routes/api.go, database/migrations/, .env

Updated (framework-managed):
  frontend/src/App.jsx, frontend/src/index.css
  frontend/vite.config.js, routes/web.go`,
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Verify we are inside an almoq3 project
		meta, err := readProjectMeta()
		if err != nil {
			fmt.Printf("%s Not inside an almoq3 project (almoq3.json not found).\n", color.RedString("❌"))
			os.Exit(1)
		}

		currentVersion := meta.FrameworkVersion
		latestVersion := Version

		// Print header
		fmt.Printf("\n")
		fmt.Printf("%s\n", color.CyanString("╔══════════════════════════════════════════════╗"))
		fmt.Printf("%s\n", color.CyanString("║       almoq3 Smart Project Upgrade           ║"))
		fmt.Printf("%s\n", color.CyanString("╚══════════════════════════════════════════════╝"))
		fmt.Printf("  Project:  %s\n", color.YellowString(meta.ProjectName))
		fmt.Printf("  Current:  %s\n", color.RedString("v"+currentVersion))
		fmt.Printf("  Latest:   %s\n\n", color.GreenString("v"+latestVersion))

		if currentVersion == latestVersion {
			fmt.Printf("%s Already up to date! Your project is on v%s\n", color.GreenString("✓"), latestVersion)
			return
		}

		fmt.Printf("%s Upgrading framework files...\n", color.CyanString("⚙️"))
		fmt.Printf("%s Your controllers, models, routes, and migrations are SAFE.\n\n", color.BlueString("🛡️"))

		// 2. Template data
		tmplData := struct {
			ProjectName      string
			FrameworkVersion string
			CreatedAt        string
		}{
			ProjectName:      meta.ProjectName,
			FrameworkVersion: latestVersion,
			CreatedAt:        meta.CreatedAt,
		}

		// 3. Framework-owned files only
		frameworkFiles := map[string]string{
			"app_jsx.tmpl":           "frontend/src/App.jsx",
			"ui_showcase_jsx.tmpl":   "frontend/src/UIShowcase.jsx",
			"main_jsx.tmpl":          "frontend/src/main.jsx",
			"globals_css.tmpl":       "frontend/src/index.css",
			"vite_config.tmpl":       "frontend/vite.config.js",
			"tailwind_config.tmpl":   "frontend/tailwind.config.js",
			"vite_package_json.tmpl": "frontend/package.json",
			"routes_web.tmpl":        "routes/web.go",
		}

		updatedCount := 0
		for tmplName, targetFile := range frameworkFiles {
			if err := upgradeFromTemplate(tmplName, targetFile, tmplData); err != nil {
				fmt.Printf("  %s Skipped %-40s %v\n", color.YellowString("⚠"), targetFile, err)
			} else {
				fmt.Printf("  %s Updated: %s\n", color.GreenString("✓"), color.WhiteString(targetFile))
				updatedCount++
			}
		}

		// 4. Rebuild frontend
		fmt.Printf("\n%s Rebuilding frontend assets...\n", color.CyanString("⚙️"))
		npmName := "npm"
		if runtime.GOOS == "windows" {
			npmName = "npm.cmd"
		}

		if _, err := os.Stat("frontend"); err == nil {
			npmInstall := exec.Command(npmName, "install")
			npmInstall.Dir = "frontend"
			if err := npmInstall.Run(); err != nil {
				fmt.Printf("  %s npm install failed. Run manually: cd frontend && npm install\n", color.YellowString("⚠️"))
			} else {
				fmt.Printf("  %s Dependencies updated.\n", color.GreenString("✓"))
			}

			npmBuild := exec.Command(npmName, "run", "build")
			npmBuild.Dir = "frontend"
			if err := npmBuild.Run(); err != nil {
				fmt.Printf("  %s Build failed. Run manually: cd frontend && npm run build\n", color.YellowString("⚠️"))
			} else {
				fmt.Printf("  %s Frontend assets rebuilt.\n", color.GreenString("✓"))
			}
		}

		// 5. Update almoq3.json version
		meta.FrameworkVersion = latestVersion
		if jsonData, err := json.MarshalIndent(meta, "", "  "); err == nil {
			os.WriteFile("almoq3.json", jsonData, 0644)
		}

		// 6. Success summary
		fmt.Printf("\n%s\n", color.GreenString("╔══════════════════════════════════════════════╗"))
		fmt.Printf("%s\n", color.GreenString("║           Upgrade Complete! 🎉               ║"))
		fmt.Printf("%s\n", color.GreenString("╚══════════════════════════════════════════════╝"))
		fmt.Printf("  Files updated:  %s\n", color.GreenString(fmt.Sprintf("%d", updatedCount)))
		fmt.Printf("  Version:        %s → %s\n", color.RedString("v"+currentVersion), color.GreenString("v"+latestVersion))
		fmt.Printf("\n  %s Restart your server: %s\n\n", color.CyanString("→"), color.YellowString("go run main.go"))
	},
}

func readProjectMeta() (*ProjectMeta, error) {
	data, err := os.ReadFile("almoq3.json")
	if err != nil {
		return nil, err
	}
	var meta ProjectMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}
	return &meta, nil
}

func upgradeFromTemplate(tmplName, targetPath string, data interface{}) error {
	content, err := templates.Files.ReadFile(tmplName)
	if err != nil {
		return fmt.Errorf("template %s not found", tmplName)
	}

	tmpl, err := template.New(tmplName).Parse(string(content))
	if err != nil {
		return fmt.Errorf("parse error: %v", err)
	}

	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return err
	}

	f, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
