package controllers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/GeorgeGogos/Company-Management/models"
	"github.com/GeorgeGogos/Company-Management/repository"
)

type UserController struct {
	UserRepo *repository.UserRepository
}

// Register a new user
func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Unable to hash password", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)

	// Save user in the database
	err = uc.UserRepo.CreateUser(&user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}
