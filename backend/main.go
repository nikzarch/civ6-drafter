package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"project/handlers"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5674"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	})) // middleware cors filter
	r.Static("/static", "./static")
	api := r.Group("/api")
	{
		api.GET("/nations", handlers.HandleNations)
		api.GET("/draft", handlers.HandleDrafts)
	}

	r.Run("localhost:8080")
}
