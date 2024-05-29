package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
	inmemoryrepo "github.com/yusufsakhtar/playstation-assignment/internal/repository/in-memory"
	"github.com/yusufsakhtar/playstation-assignment/internal/service"
	"github.com/yusufsakhtar/playstation-assignment/web/handlers"
)

func main() {
	useInMemory := true   // Toggle this to switch between in-memory and SQLite
	seedFromFiles := true // Toggle this to seed from files or not

	var inventoryRepo repository.InventoryRepo
	var userRepo repository.UserRepo
	var cartRepo repository.CartRepo
	var orderRepo repository.OrderRepo
	var userService *service.UserService
	var cartService *service.CartService
	var checkoutService *service.CheckoutService

	if useInMemory {
		inventoryRepo = inmemoryrepo.NewInMemoryInventoryRepo(seedFromFiles, "sample_input/inventory.json")
		userRepo = inmemoryrepo.NewInMemoryUserRepo(seedFromFiles, "sample_input/users.json")
		cartRepo = inmemoryrepo.NewInMemoryCartRepo(seedFromFiles, "sample_input/carts.json")
		orderRepo = inmemoryrepo.NewInMemoryOrderRepo()
		userService = service.NewUserService(userRepo, cartRepo)
		cartService = service.NewCartService(inventoryRepo, cartRepo)
		checkoutService = service.NewCheckoutService(inventoryRepo, cartRepo, orderRepo)
	} else {
		log.Fatal("SQLite not implemented yet")
	}

	mux := mux.NewRouter()
	handlers.RegisterHandlers(mux, userService, cartService, checkoutService, inventoryRepo, userRepo, cartRepo, orderRepo)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
