package routers

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func runCommandInDir(dir string, name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Command %s failed: %s", name, err)
	}
}

func createFile(path, content string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatalf("Failed to write to file: %s", err)
	}
}

func createProjectStructure(projectName string) {
	dirs := []string{
		filepath.Join(projectName, "internal/handlers"),
	}
	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatalf("Failed to create directory: %s", err)
		}
	}
}

func getModuleName(projectName string) string {
	return strings.ReplaceAll(projectName, "/", "-")
}
