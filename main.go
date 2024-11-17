package main

import (
	"database/sql"
	"fmt"
	"os"
	"quiz-3/controllers"
	"quiz-3/database"
	"quiz-3/middleware"
	"quiz-3/repository"

	_ "quiz-3/docs" // Mengimpor file docs.go yang dihasilkan oleh swag

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	DB  *sql.DB
	err error
)

// @title Category API Documentation
// @version 1.0
// @description This is the API documentation for the Category service and Books API
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url http://example.com/contact
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB)
	middleware.DB = DB
	defer middleware.DB.Close() // Pastikan koneksi ditutup saat aplikasi selesai

	router := gin.Default()

	userRepo := &repository.UserRepository{DB: DB}
	categoryRepo := &repository.CategoryRepository{DB: DB}
	bookRepo := &repository.BookRepository{DB: DB}

	// Login route
	router.POST("/api/users/login", controllers.Login(userRepo))
	router.POST("/api/users/register", controllers.RegisterUser(userRepo))

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/categories", controllers.GetAllCategories(categoryRepo))
	protected.GET("/categories/:id", controllers.GetCategoryByID(categoryRepo))
	protected.POST("/categories", controllers.CreateCategory(categoryRepo))
	protected.DELETE("/categories/:id", controllers.DeleteCategory(categoryRepo))
	protected.GET("/categories/:id/books", controllers.GetBooksByCategoryID(categoryRepo))
	protected.GET("/books", controllers.GetAllBooks(bookRepo))
	protected.GET("/books/:id", controllers.GetBookByID(bookRepo))
	protected.POST("/books", controllers.CreateBook(bookRepo))
	protected.DELETE("/books/:id", controllers.DeleteBook(bookRepo))

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	router.Run(":" + os.Getenv("PORT"))
}
