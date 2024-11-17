package controllers

import (
	"net/http"
	"quiz-3/helpers"
	"quiz-3/repository"
	"quiz-3/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllCategories godoc
// @Summary Get all categories
// @Description Retrieve a list of all categories
// @Tags Categories
// @Accept json
// @Produce json
// @Success 200 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/categories [get]
func GetAllCategories(repo *repository.CategoryRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := repo.GetAllCategories()
		if err != nil {
			helpers.ResponseJSON(c, http.StatusInternalServerError, "Failed to fetch categories", "error", nil)
			return
		}
		helpers.ResponseJSON(c, http.StatusOK, "Fetched categories successfully", "success", categories)
	}
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Description Retrieve a category by its ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /api/categories/{id} [get]
// GetCategoryByID handler for fetching category by ID
func GetCategoryByID(repo *repository.CategoryRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		category, err := repo.GetCategoryByID(id)
		if err != nil {
			helpers.ResponseJSON(c, http.StatusNotFound, "Category not found", "error", nil)
			return
		}
		helpers.ResponseJSON(c, http.StatusOK, "Fetched category successfully", "success", category)
	}
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Add a new category to the database
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body structs.Category true "Category Data"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/categories [post]
func CreateCategory(repo *repository.CategoryRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload structs.Category
		if err := c.ShouldBindJSON(&payload); err != nil || payload.Name == "" {

			helpers.ResponseJSON(c, http.StatusBadRequest, "Invalid input", "error", nil)
			return
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

		err := repo.InsertCategory(payload.Name, createdBy)
		if err != nil {
			helpers.ResponseJSON(c, http.StatusInternalServerError, "Failed to create category", "error", nil)

			return
		}
		helpers.ResponseJSON(c, http.StatusCreated, "Category created successfully", "success", nil)

	}
}

// DeleteCategory godoc
// @Summary Delete a category by ID
// @Description Remove a category by its ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /api/categories/{id} [delete]
func DeleteCategory(repo *repository.CategoryRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		err := repo.DeleteCategory(id)
		if err != nil {
			helpers.ResponseJSON(c, http.StatusNotFound, "Category not found", "error", nil)
			return
		}
		helpers.ResponseJSON(c, http.StatusOK, "Category deleted successfully", "success", nil)
	}
}

// GetBooksByCategoryID godoc
// @Summary Get books by category ID
// @Description Retrieve books that belong to a specific category
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /api/categories/{id}/books [get]
func GetBooksByCategoryID(repo *repository.CategoryRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		books, err := repo.GetBooksByCategoryID(id)
		if err != nil {
			helpers.ResponseJSON(c, http.StatusNotFound, "Category not found", "error", nil)
			return
		}
		helpers.ResponseJSON(c, http.StatusOK, "Books fetched successfully", "success", books)

	}
}
