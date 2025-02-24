package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"qunatity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost TIme", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func _getBookByID(id string) (*book, error) {
	for idx, book := range books {
		if book.ID == id {
			return &books[idx], nil
		}
	}
	return nil, errors.New("book not found")
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := _getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, book)
	}
}

func issueBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}

	book, err := _getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not available to checkout"})
		return
	}

	book.Quantity -= 1

	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}

	book, err := _getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "book added to library"})
}

func removeBook(c *gin.Context) {
	id := c.Param("id")
	_, err := _getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	var newLibrary = []book{}

	for _, book := range books {
		if id != book.ID {
			newLibrary = append(newLibrary, book)
		}
	}
	books = newLibrary
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "book has been deleted from library"})
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/book/:id", getBookByID)
	router.POST("/create", createBook)
	router.PUT("/checkout", issueBook)
	router.PUT("/return", returnBook)
	router.DELETE("/delete/:id", removeBook)

	router.Run("localhost:8080")
}
