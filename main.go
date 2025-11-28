package main

import (
	"log"
	"os"

	"practice-golang/config"
	"practice-golang/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	
	config.LoadEnv() 
	config.ConnectDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(r)

	log.Printf("Listening on :%s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
