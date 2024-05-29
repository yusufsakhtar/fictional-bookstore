package handlers

import (
	"github.com/gorilla/mux"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
	"github.com/yusufsakhtar/playstation-assignment/internal/service"
)

func RegisterHandlers(
	mux *mux.Router,
	userService *service.UserService,
	cartService *service.CartService,
	orderService *service.OrderService,
	inventoryRepo repository.InventoryRepo,
	userRepo repository.UserRepo,
	cartRepo repository.CartRepo,
	orderRepo repository.OrderRepo,
) {
	mux.HandleFunc("/users", ListUsers(userService)).Methods("GET")
	mux.HandleFunc("/users", CreateUser(userService)).Methods("POST")
	mux.HandleFunc("/users/{id}", GetUser(userService)).Methods("GET")
	mux.HandleFunc("/users/{id}", DeleteUser(userService)).Methods("DELETE")
	mux.HandleFunc("/users/{id}/cart", GetUserCart(userService)).Methods("GET")
	// TODO: change to POST? Patch is more "idiomatic" but POST is more commonly used
	// TODO: move logic to /carts/{id} (I wrote this before I went deeper into cart/order/checkout logic)
	mux.HandleFunc("/users/{id}/cart", AddItemsToUserCart(cartService)).Methods("PATCH")

	mux.HandleFunc("/carts", ListCarts(cartService)).Methods("GET")
	mux.HandleFunc("/carts/{id}", GetCart(cartService)).Methods("GET")
	mux.HandleFunc("/carts/{id}/checkout", CheckoutCart(cartService)).Methods("POST")

	mux.HandleFunc("/orders", ListOrders(orderService)).Methods("GET")
	mux.HandleFunc("/orders/{id}", GetOrder(orderService)).Methods("GET")
	mux.HandleFunc("/orders/{id}/confirm", ConfirmOrder(orderService)).Methods("POST")
	// TODO
	// mux.HandleFunc("/orders/{id}/cancel", CancelOrder(orderRepo)).Methods("POST")

	mux.HandleFunc("/inventory", ListInventory(inventoryRepo)).Methods("GET")
	mux.HandleFunc("/inventory", CreateInventoryItem(inventoryRepo)).Methods("POST")
	mux.HandleFunc("/inventory/{sku}", GetInventoryItem(inventoryRepo)).Methods("GET")
	mux.HandleFunc("/inventory/{sku}", DeleteInventoryItem(inventoryRepo)).Methods("DELETE")
	mux.HandleFunc("/inventory/{sku}", UpdateInventoryItem(inventoryRepo)).Methods("PUT")
	mux.HandleFunc("/inventory/{sku}/stock", UpdateInventoryItemStock(inventoryRepo)).Methods("PATCH")
}
