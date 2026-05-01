package controllers

import (
	"net/http"
	"strings"

	"example/go_task/config"
	"example/go_task/models"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := config.DB.Where("book_id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	if err := config.DB.Where("book_id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&models.Book{}).
		Where("book_id = ?", id).
		Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.
		Where("book_id = ?", id).
		Delete(&models.Book{}).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func searchBooks(keyword string, books []models.Book) []models.Book {
	keyword = strings.ToLower(keyword)
	var result []models.Book

	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), keyword) ||
			strings.Contains(strings.ToLower(book.Description), keyword) {
			result = append(result, book)
		}
	}

	return result
}

func concurrentSearch(keyword string, books []models.Book) []models.Book {
	chunkSize := len(books) / 4
	if chunkSize == 0 {
		chunkSize = len(books)
	}

	resultChan := make(chan []models.Book)
	var finalResults []models.Book

	for i := 0; i < len(books); i += chunkSize {
		end := i + chunkSize
		if end > len(books) {
			end = len(books)
		}

		go func(subset []models.Book) {
			resultChan <- searchBooks(keyword, subset)
		}(books[i:end])
	}

	// Collect results
	numChunks := (len(books) + chunkSize - 1) / chunkSize
	for i := 0; i < numChunks; i++ {
		res := <-resultChan
		finalResults = append(finalResults, res...)
	}

	return finalResults
}

func SearchBooks(c *gin.Context) {
	query := c.Query("q")

	if query == "" {
		c.JSON(400, gin.H{"error": "query parameter 'q' is required"})
		return
	}

	var books []models.Book
	config.DB.Find(&books)

	results := concurrentSearch(query, books)

	c.JSON(200, results)
}
