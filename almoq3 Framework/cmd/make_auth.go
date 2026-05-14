package cmd

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var makeAuthCmd = &cobra.Command{
	Use:   "make:auth",
	Short: "Scaffold a complete JWT-based authentication system",
	Run: func(cmd *cobra.Command, args []string) {
		if !IsProjectRoot() {
			color.Red("❌ Error: This command can only be executed from the root directory of an almoq3 project.")
			return
		}

		fmt.Printf("%s Scaffolding Authentication System...\n", color.CyanString("🛡️"))

		// Configuration for templates
		projectName := getProjectName()
		data := map[string]string{
			"ProjectName": projectName,
		}

		scaffoldFiles := []struct {
			tmpl   string
			target string
		}{
			{"auth_controller.tmpl", "app/controllers/auth_controller.go"},
			{"user_model.tmpl", "app/models/user.go"},
			{"auth_middleware_jwt.tmpl", "app/middleware/auth_middleware.go"},
			{"user_migration.tmpl", fmt.Sprintf("database/migrations/%d_create_users_table.go", time.Now().Unix())},
		}

		for _, f := range scaffoldFiles {
			fmt.Printf("  %s Generating %s...\n", color.YellowString("⚡"), f.target)
			if err := generateFromTemplate(f.tmpl, f.target, data); err != nil {
				fmt.Printf("  %s Failed: %v\n", color.RedString("❌"), err)
				return
			}
		}

		fmt.Printf("\n%s Authentication system scaffolded successfully!\n", color.GreenString("✅"))
		fmt.Printf("Next steps:\n")
		fmt.Printf("  1. Run `go mod tidy` to install dependencies.\n")
		fmt.Printf("  2. Run `almoq3 migrate` to create the users table.\n")
		fmt.Printf("  3. Register routes in `routes/api.go` using the new AuthController.\n")
	},
}

func init() {
	rootCmd.AddCommand(makeAuthCmd)
}
