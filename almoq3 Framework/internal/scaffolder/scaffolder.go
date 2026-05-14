package scaffolder

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"text/template"
	"time"

	"github.com/almoq3/almoq3-cli/internal/templates"
	"github.com/fatih/color"
)

type Config struct {
	ProjectName string
	CreatedAt   string
}

func Scaffold(projectName string) error {
	config := Config{
		ProjectName: projectName,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	fmt.Printf("%s Creating project: %s\n", color.CyanString("⚙️"), color.YellowString(projectName))

	// 1. Create directory structure
	dirs := []string{
		"app/controllers",
		"app/models",
		"app/services",
		"app/providers",
		"app/middleware",
		"bootstrap",
		"config",
		"database/migrations",
		"database/seeders",
		"routes",
		"storage/logs",
		"storage/uploads",
		"storage/framework/cache",
		"cmd/seed",
		"app/requests",
		"resources/views",
		"frontend/src/components",
		"frontend/src/lib",
		"public",
	}

	for _, dir := range dirs {
		path := filepath.Join(projectName, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
		fmt.Printf("  %s Created directory: %s\n", color.GreenString("✓"), dir)
	}

	// 2. Generate files from templates
	files := map[string]string{
		"env.tmpl":               ".env",
		"main.tmpl":              "main.go",
		"makefile.tmpl":          "Makefile",
		"gitignore.tmpl":         ".gitignore",
		"go_mod.tmpl":            "go.mod",
		"config_app.tmpl":        "config/app.go",
		"config_database.tmpl":   "config/database.go",
		"config_cache.tmpl":      "config/cache.go",
		"config_master.tmpl":     "config/config.go",
		"bootstrap_config.tmpl":  "bootstrap/config.go",
		"bootstrap_database.tmpl": "bootstrap/database.go",
		"bootstrap_cache.tmpl":    "bootstrap/cache.go",
		"bootstrap_logger.tmpl":   "bootstrap/logger.go",
		"routes_api.tmpl":        "routes/api.go",
		"routes_web.tmpl":        "routes/web.go",
		"routes_register.tmpl":   "routes/routes.go",
		"database_seeder.tmpl":   "database/seeders/database_seeder.go",
		"seed_main.tmpl":         "cmd/seed/main.go",
		"dockerfile.tmpl":        "Dockerfile",
		"docker_compose.tmpl":    "docker-compose.yml",
		"package_json.tmpl":      "package.json",
		"almoq3_json.tmpl":       "almoq3.json",
		"proxy.tmpl":             "artisan.bat",
		"vite_package_json.tmpl": "frontend/package.json",
		"vite_config.tmpl":       "frontend/vite.config.js",
		"tailwind_config.tmpl":   "frontend/tailwind.config.js",
		"globals_css.tmpl":       "frontend/src/index.css",
		"app_jsx.tmpl":           "frontend/src/App.jsx",
		"ui_showcase_jsx.tmpl":   "frontend/src/UIShowcase.jsx",
		"main_jsx.tmpl":          "frontend/src/main.jsx",
	}

	for tmplName, fileName := range files {
		if err := generateFile(projectName, tmplName, fileName, config); err != nil {
			return err
		}
		fmt.Printf("  %s Generated file: %s\n", color.GreenString("✓"), fileName)
	}

	// Create frontend/index.html
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
	os.WriteFile(filepath.Join(projectName, "frontend/index.html"), []byte(indexHTML), 0644)

	fmt.Printf("\n%s Setting up frontend environment (this may take a minute)...\n", color.CyanString("⚙️"))

	npmName := "npm"
	if runtime.GOOS == "windows" {
		npmName = "npm.cmd"
	}

	// Run npm install in frontend directory
	npmInstall := exec.Command(npmName, "install")
	npmInstall.Dir = filepath.Join(projectName, "frontend")
	if err := npmInstall.Run(); err != nil {
		fmt.Printf("%s Warning: Failed to run npm install. Please run it manually in the frontend directory.\n", color.YellowString("⚠️"))
	} else {
		fmt.Printf("  %s Frontend dependencies installed.\n", color.GreenString("✓"))
	}

	// Run npm run build
	npmBuild := exec.Command(npmName, "run", "build")
	npmBuild.Dir = filepath.Join(projectName, "frontend")
	if err := npmBuild.Run(); err != nil {
		fmt.Printf("%s Warning: Failed to run npm run build.\n", color.YellowString("⚠️"))
	} else {
		fmt.Printf("  %s Frontend assets built.\n", color.GreenString("✓"))
	}

	fmt.Printf("\n%s Project %s created successfully!\n", color.CyanString("✨"), color.YellowString(projectName))
	fmt.Printf("Next steps:\n")
	fmt.Printf("  cd %s\n", projectName)
	fmt.Printf("  go mod tidy\n")
	fmt.Printf("  go run main.go %s\n", color.HiBlackString("(or make run)"))

	return nil
}

func generateFile(projectName, tmplName, fileName string, config Config) error {
	data, err := templates.Files.ReadFile(tmplName)
	if err != nil {
		return fmt.Errorf("failed to read template %s: %w", tmplName, err)
	}

	tmpl, err := template.New(tmplName).Parse(string(data))
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", tmplName, err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, config); err != nil {
		return fmt.Errorf("failed to execute template %s: %w", tmplName, err)
	}

	outputPath := filepath.Join(projectName, fileName)
	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", outputPath, err)
	}

	return nil
}
