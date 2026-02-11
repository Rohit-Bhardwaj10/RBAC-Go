package main

import (
	"fmt"
	"log"

	"github.com/Rohit-Bhardwaj10/RBAC-Go/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load("cmd/api/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func loadDatabase() {
	db.InitDb()
}

func serveApplication() {
	router := gin.Default()
	router.Run(":8080")
	fmt.Println("Server running on port 8080")
}

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
	fmt.Println("Application started successfully")
}
