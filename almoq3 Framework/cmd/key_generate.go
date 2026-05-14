package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var keyGenerateCmd = &cobra.Command{
	Use:   "key:generate",
	Short: "Set the application key",
	Run: func(cmd *cobra.Command, args []string) {
		if !IsProjectRoot() {
			color.Red("❌ Error: This command can only be executed from the root directory of an almoq3 project.")
			return
		}

		key := generateRandomKey(32)
		fmt.Printf("%s Generating Application Key...\n", color.CyanString("⚡"))

		err := updateEnvKey("APP_KEY", key)
		if err != nil {
			fmt.Printf("%s Failed to update .env: %v\n", color.RedString("❌"), err)
			return
		}

		fmt.Printf("%s Application key [%s] set successfully.\n", color.GreenString("✅"), color.YellowString(key))
	},
}

func generateRandomKey(length int) string {
	b := make([]byte, length/2)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func updateEnvKey(key, value string) error {
	input, err := os.ReadFile(".env")
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")
	found := false
	for i, line := range lines {
		if strings.HasPrefix(line, key+"=") {
			lines[i] = fmt.Sprintf("%s=%s", key, value)
			found = true
			break
		}
	}

	if !found {
		lines = append(lines, fmt.Sprintf("%s=%s", key, value))
	}

	output := strings.Join(lines, "\n")
	return os.WriteFile(".env", []byte(output), 0644)
}

func init() {
	rootCmd.AddCommand(keyGenerateCmd)
}
