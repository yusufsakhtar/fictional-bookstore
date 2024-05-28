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
	var userService *service.UserService
	var cartService *service.CartService

	if useInMemory {
		inventoryRepo = inmemoryrepo.NewInMemoryInventoryRepo(seedFromFiles, "sample_input/inventory.json")
		userRepo = inmemoryrepo.NewInMemoryUserRepo(seedFromFiles, "sample_input/users.json")
		cartRepo = inmemoryrepo.NewInMemoryCartRepo()
		userService = service.NewUserService(userRepo, cartRepo)
		cartService = service.NewCartService(inventoryRepo, cartRepo)
	} else {
		log.Fatal("SQLite not implemented yet")
	}

	mux := mux.NewRouter()
	handlers.RegisterHandlers(mux, userService, cartService, inventoryRepo, userRepo, cartRepo)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
