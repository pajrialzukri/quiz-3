package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"quiz-3/structs"
)

type BookRepository struct {
	DB *sql.DB
}

// GetAllBooks - Mendapatkan semua buku
func (r *BookRepository) GetAllBooks() ([]structs.Book, error) {
	rows, err := r.DB.Query("SELECT * FROM books")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var books []structs.Book
	for rows.Next() {
		var book structs.Book
		err := rows.Scan(
			&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear,
			&book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID,
			&book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// GetBookByID - Mendapatkan detail buku
func (r *BookRepository) GetBookByID(id int) (*structs.Book, error) {
	var book structs.Book
	err := r.DB.QueryRow("SELECT * FROM books WHERE id = $1", id).Scan(
		&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear,
		&book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID,
		&book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("book not found")
	}
	return &book, err
}

// InsertBook - Menambahkan buku baru
func (r *BookRepository) InsertBook(book structs.Book) error {
	query := `
		INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_by, modified_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := r.DB.Exec(query,
		book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.CreatedBy, book.CreatedBy,
	)
	if err != nil {
		fmt.Println("InsertBook error:", err)
	}
	return err
}

// DeleteBook - Menghapus buku
func (r *BookRepository) DeleteBook(id int) error {
	result, err := r.DB.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("book not found")
	}
	return nil
}
