package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books = []Book{
    {ID: 1, Title: "Learn Go", Author: "Author A"},
    {ID: 2, Title: "Gin for Beginners", Author: "Author B"},
}

func main() {
    router := gin.Default()

    router.GET("/books", getBooks)
    router.GET("/books/:id", getBook)
    router.POST("/books", addBook)
    router.PUT("/books/:id", updateBook)
    router.DELETE("/books/:id", deleteBook)

    router.Run(":8080")
}

func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
        return
    }

    for _, book := range books {
        if book.ID == id {
            c.IndentedJSON(http.StatusOK, book)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func addBook(c *gin.Context) {
    var newBook Book

    if err := c.BindJSON(&newBook); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
        return
    }

    newBook.ID = len(books) + 1
    books = append(books, newBook)
    c.IndentedJSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
        return
    }

    var updatedBook Book
    if err := c.BindJSON(&updatedBook); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
        return
    }

    for i, book := range books {
        if book.ID == id {
            books[i] = updatedBook
            books[i].ID = id
            c.IndentedJSON(http.StatusOK, books[i])
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func deleteBook(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
        return
    }

    for i, book := range books {
        if book.ID == id {
            books = append(books[:i], books[i+1:]...)
            c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted"})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}
