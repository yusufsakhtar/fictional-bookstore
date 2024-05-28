package handlers

import (
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"

	"encoding/json"
	"net/http"
)

// GetUser handles GET requests to retrieve a user by ID.
func GetUser(repo repository.UserRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		user, err := repo.GetUser(repository.GetUserInput{ID: id})
		if err != nil || user == nil {
			http.Error(w, "Unable to retrieve user", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

// CreateUser handles POST requests to create a new user.
func CreateUser(repo repository.UserRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createUserInput repository.CreateUserInput
		if err := json.NewDecoder(r.Body).Decode(&createUserInput); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		if err := repo.CreateUser(createUserInput); err != nil {
			http.Error(w, "Unable to create user", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// DeleteUser handles DELETE requests to delete a user by ID.
func DeleteUser(repo repository.UserRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if err := repo.DeleteUser(repository.DeleteUserInput{ID: id}); err != nil {
			http.Error(w, "Unable to delete user", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

// ListUsers handles GET requests to list all users.
func ListUsers(repo repository.UserRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := repo.ListUsers()
		if err != nil {
			http.Error(w, "Unable to list users", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}
