package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/earlysvahn/kickstart-go-api/internal/routers"
	"github.com/earlysvahn/kickstart-go-api/internal/utils"
	"github.com/manifoldco/promptui"
)

var version = "1.0.6"

func promptUser(prompt string) string {
	promptTemplate := promptui.Prompt{
		Label: prompt,
	}
	response, err := promptTemplate.Run()
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	return strings.TrimSpace(response)
}

func selectRouter() string {
	prompt := promptui.Select{
		Label: "Choose router",
		Items: []string{"gin", "mux", "chi"},
	}

	_, router, err := prompt.Run()
	if err != nil {
		log.Fatalf("Failed to select router: %v", err)
	}

	return router
}

func main() {
	showVersion := flag.Bool("version", false, "Show the current version of the tool")
	var router string
	flag.StringVar(&router, "router", "", "Choose router for your Go API project: gin, mux, or chi")
	flag.Parse()

	if *showVersion {
		updateAvailable, latestVersion := utils.CheckForUpdates(version)
		fmt.Printf("Kickstart Go API version: %s\n", version)

		if updateAvailable {
			fmt.Printf("A new version is available: %s\n", latestVersion)
			fmt.Println("Update your tool with:")
			fmt.Println("go install github.com/earlysvahn/kickstart-go-api@latest")
		}
		return
	}

	var projectName string
	if len(flag.Args()) > 0 {
		projectName = flag.Args()[0]
	} else {
		projectName = promptUser("Please enter the project name")
	}

	if router == "" {
		router = selectRouter()
	}

	router = strings.ToLower(router)
	switch router {
	case "gin":
		fmt.Println("Creating project with Gin router")
		routers.CreateGinProject(projectName)
	case "chi":
		fmt.Println("Creating project with Chi router")
		routers.CreateChiProject(projectName)
	case "mux":
		fmt.Println("Creating project with Mux router")
		routers.CreateMuxProject(projectName)
	default:
		log.Fatalf("Unsupported router: %s. Use gin, mux, or chi.", router)
	}

	fmt.Printf("Project %s has been created with %s router!\n", projectName, router)
}
