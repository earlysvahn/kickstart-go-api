package routers

import (
	"fmt"
	"path/filepath"
)

func CreateMuxProject(projectName string) {
	createProjectStructure(projectName)
	runCommandInDir(projectName, "go", "mod", "init", projectName)
	runCommandInDir(projectName, "go", "get", "-u", "github.com/gorilla/mux")

	moduleName := getModuleName(projectName)
	mainFile := filepath.Join(projectName, "main.go")
	createFile(mainFile, muxMainGoContent(moduleName))

	handlersFile := filepath.Join(projectName, "internal/handlers/handlers.go")
	createFile(handlersFile, muxHandlersGoContent())

	envFile := filepath.Join(projectName, ".env")
	createFile(envFile, "PORT=8080\n")
}

func muxMainGoContent(moduleName string) string {
	return fmt.Sprintf(`package main

import (
	"log"
	"net/http"
	"os"
	"%s/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Ping).Methods("GET")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":" + port, r))
}`, moduleName)
}

func muxHandlersGoContent() string {
	return `package handlers

import "net/http"
import "encoding/json"

func Ping(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello from mux!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}`
}
