package handlers

import (
	"net/http"

	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
)

func RegisterHandlers(mux *http.ServeMux, inventoryRepo repository.InventoryRepo, userRepo repository.UserRepo) {
	// mux.HandleFunc("/items", ViewItems(itemRepo))
	// mux.HandleFunc("/item", GetItem(itemRepo))
	// mux.HandleFunc("/item/create", CreateItem(itemRepo))
	// mux.HandleFunc("/item/update", UpdateItem(itemRepo))
	// mux.HandleFunc("/item/delete", DeleteItem(itemRepo))

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ListUsers(userRepo)(w, r)
		case http.MethodPost:
			CreateUser(userRepo)(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetUser(userRepo)(w, r)
			// TODO
		// case http.MethodPut:
		case http.MethodDelete:
			DeleteUser(userRepo)(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// mux.HandleFunc("/cart", ViewCart(cartRepo))
	// mux.HandleFunc("/cart/add", AddItemToCart(cartRepo))
	// mux.HandleFunc("/cart/remove", RemoveItemFromCart(cartRepo))
}
