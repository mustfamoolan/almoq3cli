package scaffolder

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
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
		"routes_api.tmpl":       "routes/api.go",
		"routes_web.tmpl":       "routes/web.go",
		"routes_register.tmpl":  "routes/routes.go",
		"proxy.tmpl":            "artisan.bat",
		"database_seeder.tmpl":  "database/seeders/database_seeder.go",
		"seed_main.tmpl":        "cmd/seed/main.go",
		"bootstrap_logger.tmpl":  "bootstrap/logger.go",
		"bootstrap_cache.tmpl":   "bootstrap/cache.go",
		"dockerfile.tmpl":       "Dockerfile",
		"docker_compose.tmpl":   "docker-compose.yml",
		"package_json.tmpl":     "package.json",
		"almoq3_json.tmpl":      "almoq3.json",
		"welcome.tmpl":          "resources/views/welcome.html",
	}

	for tmplName, fileName := range files {
		if err := generateFile(projectName, tmplName, fileName, config); err != nil {
			return err
		}
		fmt.Printf("  %s Generated file: %s\n", color.GreenString("✓"), fileName)
	}

	fmt.Printf("\n%s Project %s created successfully!\n", color.CyanString("✨"), color.YellowString(projectName))
	fmt.Printf("Next steps:\n")
	fmt.Printf("  cd %s\n", projectName)
	fmt.Printf("  go mod tidy\n")
	fmt.Printf("  make run\n")

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
