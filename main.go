package main

import (
	"example/go_task/config"
	"example/go_task/models"

	"example/go_task/controllers"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	config.ConnectDB()
	config.DB.AutoMigrate(&models.Book{})

	router := gin.Default()

	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBook)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
	router.GET("/books/search", controllers.SearchBooks)

	router.Run(":8080")

}
