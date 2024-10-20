package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/earlysvahn/kickstart-go-api/internal/routers"
)

func main() {
	var router string
	flag.StringVar(&router, "router", "mux", "Choose router for your Go API project: gin, mux, or chi")

	if len(os.Args) >= 4 {
		err := flag.CommandLine.Set("router", os.Args[3])
		if err != nil {
			log.Fatalf("Failed to set router flag: %v", err)
		}
	}

	flag.Parse()

	fmt.Printf("Router flag: %s\n", router)

	if len(flag.Args()) == 0 {
		log.Fatal("Please provide a project name")
	}
	projectName := flag.Args()[0]

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
