package handlers

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
	"github.com/yusufsakhtar/playstation-assignment/internal/service"

	"encoding/json"
	"net/http"
)

// GetUser handles GET requests to retrieve a user by ID.
func GetUser(usersvc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		user, err := usersvc.GetUser(repository.GetUserInput{ID: id})
		if err != nil || user == nil {
			http.Error(w, "Unable to retrieve user", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

// CreateUser handles POST requests to create a new user.
func CreateUser(usersvc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createUserInput repository.CreateUserInput
		if err := json.NewDecoder(r.Body).Decode(&createUserInput); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		if err := usersvc.CreateUser(createUserInput); err != nil {
			http.Error(w, "Unable to create user", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// DeleteUser handles DELETE requests to delete a user by ID.
func DeleteUser(usersvc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if err := usersvc.DeleteUser(repository.DeleteUserInput{ID: id}); err != nil {
			http.Error(w, "Unable to delete user", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

// ListUsers handles GET requests to list all users.
func ListUsers(usersvc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := usersvc.ListUsers()
		if err != nil {
			http.Error(w, "Unable to list users", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

// AddItemsToUserCart handles PATCH requests to add items to a user's cart.
func AddItemsToUserCart(cartsvc *service.CartService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId := vars["id"]
		var itemIds []string
		if err := json.NewDecoder(r.Body).Decode(&itemIds); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		response, err := cartsvc.AddItemsToUserCart(repository.AddItemsToUserCartInput{UserID: userId, SKUs: itemIds})
		if err != nil {
			http.Error(w, "Unable to add items to user cart", http.StatusInternalServerError)
			return
		}

		fmt.Printf("response: %+v\n", response)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// GetUserCart handles GET requests to retrieve a user's cart by ID.
func GetUserCart(usersvc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId := vars["id"]
		cart, err := usersvc.GetUserCart(repository.GetUserCartInput{UserID: userId})
		if err != nil || cart == nil {
			http.Error(w, "Unable to retrieve user cart", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cart)
	}
}
