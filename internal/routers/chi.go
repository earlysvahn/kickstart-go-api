package routers

import (
	"fmt"
	"path/filepath"
)

func CreateChiProject(projectName string) {
	createProjectStructure(projectName)

	runCommandInDir(projectName, "go", "mod", "init", projectName)
	runCommandInDir(projectName, "go", "get", "-u", "github.com/go-chi/chi/v5")

	moduleName := getModuleName(projectName)
	mainFile := filepath.Join(projectName, "main.go")
	createFile(mainFile, chiMainGoContent(moduleName))

	handlersFile := filepath.Join(projectName, "internal/handlers/handlers.go")
	createFile(handlersFile, chiHandlersGoContent())

	envFile := filepath.Join(projectName, ".env")
	createFile(envFile, "PORT=8080\n")
}

func chiMainGoContent(moduleName string) string {
	return fmt.Sprintf(`package main

import (
	"log"
	"net/http"
	"os"
	"%s/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/ping", handlers.Ping)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":" + port))
}`, moduleName)
}

func chiHandlersGoContent() string {
	return `package handlers

import "net/http"
import "encoding/json"

func Ping(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "pong"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}`
}
