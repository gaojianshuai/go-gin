package main

import (
	"go-gin/config"
	"go-gin/routes"
	"go-gin/utils"
	"os"
)

func main() {
	utils.InitLogger()
	config.InitDB()

	router := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	utils.InfoLogger.Println("Starting blog system server on port " + port)
	if err := router.Run(":" + port); err != nil {
		utils.ErrorLogger.Fatal("Failed to start server:", err)
	}
}
