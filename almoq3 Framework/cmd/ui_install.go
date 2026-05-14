package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var uiInstallCmd = &cobra.Command{
	Use:   "ui:install",
	Short: "Initialize shadcn/ui + React + Vite frontend environment",
	Run: func(cmd *cobra.Command, args []string) {
		if !IsProjectRoot() {
			color.Red("❌ Error: This command can only be executed from the root directory of an almoq3 project.")
			return
		}

		fmt.Printf("%s Injecting shadcn/ui components...\n", color.CyanString("🎨"))

		projectName := getProjectName()
		data := map[string]string{
			"ProjectName": projectName,
		}

		// Create frontend directory structure
		dirs := []string{
			"frontend/src/components",
			"frontend/src/lib",
			"public",
		}
		for _, dir := range dirs {
			os.MkdirAll(dir, 0755)
		}

		// Generate Frontend Files
		files := map[string]string{
			"vite_package_json.tmpl": "frontend/package.json",
			"vite_config.tmpl":       "frontend/vite.config.js",
			"tailwind_config.tmpl":   "frontend/tailwind.config.js",
			"globals_css.tmpl":       "frontend/src/index.css",
			"app_jsx.tmpl":           "frontend/src/App.jsx",
			"ui_showcase_jsx.tmpl":   "frontend/src/UIShowcase.jsx",
			"main_jsx.tmpl":          "frontend/src/main.jsx",
		}

		for tmpl, target := range files {
			if err := generateFromTemplate(tmpl, target, data); err != nil {
				fmt.Printf("%s Failed to generate %s: %v\n", color.RedString("❌"), target, err)
				return
			}
		}

		// --- NEW: Update existing project routes to support React ---
		fmt.Printf("%s Syncing Go routes with React engine...\n", color.BlueString("🔄"))
		webRoutesPath := "routes/web.go"
		if _, err := os.Stat(webRoutesPath); err == nil {
			// Update the file to include React serving logic
			if err := generateFromTemplate("routes_web.tmpl", webRoutesPath, data); err != nil {
				fmt.Printf("%s Failed to update routes: %v\n", color.RedString("❌"), err)
			} else {
				fmt.Printf("%s Updated routes/web.go successfully.\n", color.GreenString("✓"))
			}
		}

		// Create a basic index.html for Vite
		indexHTML := `<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <link rel="icon" type="image/svg+xml" href="/vite.svg" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>almoq3 + React</title>
  </head>
  <body>
    <div id="root"></div>
    <script type="module" src="/src/main.jsx"></script>
  </body>
</html>`
		os.WriteFile("frontend/index.html", []byte(indexHTML), 0644)

		fmt.Printf("%s Frontend engine ready!\n", color.GreenString("🚀"))
		fmt.Println("\nNext steps:")
		color.Yellow("  cd frontend")
		color.Yellow("  npm install")
		color.Yellow("  npm run dev")
		fmt.Println("\nAfter building (npm run build), your Go server will serve the frontend automatically!")
	},
}

func init() {
	rootCmd.AddCommand(uiInstallCmd)
}
