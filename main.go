package main

import (
	"GolangLearning/database"
	"GolangLearning/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.ConnectDB()

	router := gin.Default()

	router.GET("/books", handlers.GetBooks)
	router.GET("/book/:id", handlers.GetBookByID)
	router.POST("/create", handlers.CreateBook)
	router.PUT("/checkout", handlers.IssueBook)
	router.PUT("/return", handlers.ReturnBook)
	router.DELETE("/delete/:id", handlers.RemoveBook)

	router.Run("localhost:8080")
}
