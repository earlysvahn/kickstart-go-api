package routers

import (
	"fmt"
	"path/filepath"
)

func CreateGinProject(projectName string) {
	createProjectStructure(projectName)

	runCommandInDir(projectName, "go", "mod", "init", projectName)
	runCommandInDir(projectName, "go", "get", "-u", "github.com/gin-gonic/gin")

	moduleName := getModuleName(projectName)
	mainFile := filepath.Join(projectName, "main.go")
	createFile(mainFile, ginMainGoContent(moduleName))

	handlersFile := filepath.Join(projectName, "internal/handlers/handlers.go")
	createFile(handlersFile, ginHandlersGoContent())

	envFile := filepath.Join(projectName, ".env")
	createFile(envFile, "PORT=8080\n")
}

func ginMainGoContent(moduleName string) string {
	return fmt.Sprintf(`package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"log"
	"%s/internal/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/", handlers.Ping)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(router.Run(":" + port))
}`, moduleName)
}

func ginHandlersGoContent() string {
	return `package handlers

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello from gin!"})
}`
}
