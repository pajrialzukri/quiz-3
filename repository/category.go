package repository

import (
	"database/sql"
	"errors"
	"quiz-3/structs"
)

// CategoryRepository struct to handle category data
type CategoryRepository struct {
	DB *sql.DB
}

// GetAllCategories fetches all categories
func (r *CategoryRepository) GetAllCategories() ([]structs.Category, error) {
	rows, err := r.DB.Query("SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []structs.Category
	for rows.Next() {
		var category structs.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

// GetCategoryByID fetches a category by ID
func (r *CategoryRepository) GetCategoryByID(id int) (*structs.Category, error) {
	var category structs.Category
	err := r.DB.QueryRow("SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories WHERE id = $1", id).
		Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
	if err == sql.ErrNoRows {
		return nil, errors.New("category not found")
	}
	return &category, err
}

// InsertCategory inserts a new category into the database
func (r *CategoryRepository) InsertCategory(name, createdBy string) error {
	_, err := r.DB.Exec("INSERT INTO categories (name, created_at, created_by, modified_at, modified_by) VALUES ($1, CURRENT_TIMESTAMP, $2, CURRENT_TIMESTAMP, $2)", name, createdBy)
	return err
}

// DeleteCategory deletes a category by ID
func (r *CategoryRepository) DeleteCategory(id int) error {
	res, err := r.DB.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("category not found")
	}
	return nil
}

// GetBooksByCategoryID fetches all books associated with a given category ID
func (r *CategoryRepository) GetBooksByCategoryID(categoryID int) ([]structs.Book, error) {
	rows, err := r.DB.Query("SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []structs.Book
	for rows.Next() {
		var book structs.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
