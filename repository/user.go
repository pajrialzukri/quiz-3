package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"quiz-3/structs"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) GetUserByUsername(username string) (*structs.User, error) {
	var user structs.User
	query := "SELECT id, username, password FROM users WHERE username = $1"
	err := r.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) VerifyPassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func (r *UserRepository) InsertUser(username, hashedPassword string) error {
	query := `
		INSERT INTO users (username, password, created_at, created_by)
		VALUES ($1, $2, CURRENT_TIMESTAMP, $3)
	`
	_, err := r.DB.Exec(query, username, hashedPassword, "system")
	if err != nil {
		// Periksa apakah error berasal dari duplikasi username
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return fmt.Errorf("username already exists")
		}
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByID(id int) (*structs.User, error) {
	var user structs.User
	query := "SELECT id, username, password FROM users WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
