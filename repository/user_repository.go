package repository

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/GeorgeGogos/Company-Management/models"
)

type UserRepository struct {
	DB *sql.DB
}

// Create a new user
func (r *UserRepository) CreateUser(user *models.User) error {
	user.ID = uuid.New().String()
	query := `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`
	_, err := r.DB.Exec(query, user.ID, user.Email, user.Password)
	return err
}

// Find a user by email
func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, password FROM users WHERE email=$1`
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return user, err
}
