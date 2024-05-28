package handlers

import (
	"github.com/gorilla/mux"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
	"github.com/yusufsakhtar/playstation-assignment/internal/service"
)

func RegisterHandlers(mux *mux.Router, userService *service.UserService, cartService *service.CartService, inventoryRepo repository.InventoryRepo, userRepo repository.UserRepo, cartRepo repository.CartRepo) {
	mux.HandleFunc("/users", ListUsers(userService)).Methods("GET")
	mux.HandleFunc("/users", CreateUser(userService)).Methods("POST")
	mux.HandleFunc("/users/{id}", GetUser(userService)).Methods("GET")
	mux.HandleFunc("/users/{id}", DeleteUser(userService)).Methods("DELETE")
	mux.HandleFunc("/users/{id}/cart", GetUserCart(userService)).Methods("GET")
	mux.HandleFunc("/users/{id}/cart", AddItemsToUserCart(cartService)).Methods("PATCH")

	mux.HandleFunc("/inventory", ListInventory(inventoryRepo)).Methods("GET")
	mux.HandleFunc("/inventory", CreateInventoryItem(inventoryRepo)).Methods("POST")
	mux.HandleFunc("/inventory/{sku}", GetInventoryItem(inventoryRepo)).Methods("GET")
	mux.HandleFunc("/inventory/{sku}", DeleteInventoryItem(inventoryRepo)).Methods("DELETE")
	mux.HandleFunc("/inventory/{sku}", UpdateInventoryItem(inventoryRepo)).Methods("PUT")
	mux.HandleFunc("/inventory/{sku}/stock", UpdateInventoryItemStock(inventoryRepo)).Methods("PATCH")
}
