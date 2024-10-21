package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Release struct {
	TagName string `json:"tag_name"`
}

func CheckForUpdates(currentVersion string) (bool, string) {
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("https://api.github.com/repos/earlysvahn/kickstart-go-api/releases/latest")
	if err != nil {
		fmt.Println("Failed to check for updates:", err)
		return false, ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to fetch the latest version info")
		return false, ""
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		fmt.Println("Error decoding release data:", err)
		return false, ""
	}

	if release.TagName != currentVersion {
		return true, release.TagName
	}

	return false, ""
}
