package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// RepoURL is now defined in root.go

type GitHubRelease struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

var selfUpdateCmd = &cobra.Command{
	Use:   "self-update",
	Short: "Update almoq3 CLI to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s Checking for updates...\n", color.CyanString("☁️"))

		resp, err := http.Get(RepoURL)
		if err != nil {
			fmt.Printf("%s Failed to check for updates: %v\n", color.RedString("❌"), err)
			return
		}
		defer resp.Body.Close()

		var release GitHubRelease
		if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
			fmt.Printf("%s Failed to parse update info: %v\n", color.RedString("❌"), err)
			return
		}

		if release.TagName == "v"+Version {
			fmt.Printf("%s You are already using the latest version (%s).\n", color.GreenString("✅"), color.YellowString(Version))
			return
		}

		fmt.Printf("%s New version found: %s. Downloading...\n", color.CyanString("📦"), color.YellowString(release.TagName))

		// Find asset for current OS/Arch
		// For simplicity, assuming naming convention: almoq3-windows-amd64.exe
		assetName := fmt.Sprintf("almoq3-%s-%s", runtime.GOOS, runtime.GOARCH)
		if runtime.GOOS == "windows" {
			assetName += ".exe"
		}

		var downloadURL string
		for _, asset := range release.Assets {
			if asset.Name == assetName {
				downloadURL = asset.BrowserDownloadURL
				break
			}
		}

		if downloadURL == "" {
			fmt.Printf("%s Could not find a suitable binary for your OS/Arch in this release.\n", color.RedString("❌"))
			return
		}

		if err := downloadAndReplace(downloadURL); err != nil {
			fmt.Printf("%s Update failed: %v\n", color.RedString("❌"), err)
			return
		}

		fmt.Printf("%s Update successful! You are now on %s\n", color.GreenString("✅"), color.YellowString(release.TagName))
	},
}

func downloadAndReplace(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	exePath, err := os.Executable()
	if err != nil {
		return err
	}

	oldPath := exePath + ".old"
	os.Rename(exePath, oldPath)

	out, err := os.Create(exePath)
	if err != nil {
		os.Rename(oldPath, exePath) // Rollback
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		os.Rename(oldPath, exePath) // Rollback
		return err
	}

	os.Remove(oldPath)
	return nil
}

func init() {
	rootCmd.AddCommand(selfUpdateCmd)
}
