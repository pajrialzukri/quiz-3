package controllers

import (
	"net/http"
	"quiz-3/helpers"
	"quiz-3/repository"
	"quiz-3/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllBooks godoc
// @Summary Get all books
// @Description Fetch all books from the database
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {object} helpers.Response{Data=[]structs.Book} "Fetched books successfully"
// @Failure 500 {object} helpers.Response "Failed to fetch books"
// @Router /api/books [get]
func GetAllBooks(repo *repository.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := repo.GetAllBooks()
		if err != nil {
			helpers.ResponseJSON(c, http.StatusInternalServerError, "Failed to fetch books", "error", nil)
			return
		}
		helpers.ResponseJSON(c, http.StatusOK, "Fetched books successfully", "success", books)
	}
}

// GetBookByID godoc
// @Summary Get a book by ID
// @Description Fetch a book by its ID from the database
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} helpers.Response{Data=structs.Book} "Fetched book successfully"
// @Failure 404 {object} helpers.Response "Book not found"
// @Router /api/books/{id} [get]
func GetBookByID(repo *repository.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		book, err := repo.GetBookByID(id)
		if err != nil {
			helpers.ResponseJSON(c, http.StatusNotFound, "Book not found", "error", nil)
			return
		}
		helpers.ResponseJSON(c, http.StatusOK, "Fetched book successfully", "success", book)
	}
}

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book in the database
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body structs.Book true "Book data"
// @Success 201 {object} helpers.Response "Book created successfully"
// @Failure 400 {object} helpers.Response "Invalid input"
// @Failure 401 {object} helpers.Response "User not authenticated"
// @Failure 500 {object} helpers.Response "Failed to create book"
// @Router /api/books [post]
func CreateBook(repo *repository.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book structs.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			helpers.ResponseJSON(c, http.StatusBadRequest, "Invalid input", "error", nil)
			return
		}

		if book.TotalPage > 100 {
			book.Thickness = "tebal"
		} else {
			book.Thickness = "tipis"
		}

		user, exists := c.Get("user")
		if !exists {
			helpers.ResponseJSON(c, http.StatusUnauthorized, "User not authenticated", "error", nil)

			return
		}

		createdBy, ok := user.(string)
		if !ok {
			helpers.ResponseJSON(c, http.StatusUnauthorized, "Failed to extract user information", "error", nil)

			return
		}

		book.CreatedBy = createdBy

		err := repo.InsertBook(book)
		if err != nil {
			helpers.ResponseJSON(c, http.StatusInternalServerError, "Failed to create book", "error", nil)
			return
		}
		helpers.ResponseJSON(c, http.StatusCreated, "Book created successfully", "success", nil)
	}
}

// DeleteBook godoc
// @Summary Delete a book by ID
// @Description Delete a book from the database by its ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} helpers.Response "Book deleted successfully"
// @Failure 404 {object} helpers.Response "Book not found"
// @Router /api/books/{id} [delete]
func DeleteBook(repo *repository.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		err := repo.DeleteBook(id)
		if err != nil {
			helpers.ResponseJSON(c, http.StatusNotFound, err.Error(), "error", nil)
			return
		}
		helpers.ResponseJSON(c, http.StatusOK, "Book deleted successfully", "success", nil)
	}
}
