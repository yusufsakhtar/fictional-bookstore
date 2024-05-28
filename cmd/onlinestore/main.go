package main

import (
	"log"
	"net/http"

	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
	inmemoryrepo "github.com/yusufsakhtar/playstation-assignment/internal/repository/in-memory"
	"github.com/yusufsakhtar/playstation-assignment/web/handlers"
)

func main() {
	useInMemory := true   // Toggle this to switch between in-memory and SQLite
	seedFromFiles := true // Toggle this to seed from files or not

	var inventoryRepo repository.InventoryRepo
	var userRepo repository.UserRepo

	if useInMemory {
		inventoryRepo = inmemoryrepo.NewInMemoryInventoryRepo(seedFromFiles, "sample_input/inventory.json")
		userRepo = inmemoryrepo.NewInMemoryUserRepo(seedFromFiles, "sample_input/users.json")
	} else {
		log.Fatal("SQLite not implemented yet")
	}

	mux := http.NewServeMux()
	handlers.RegisterHandlers(mux, inventoryRepo, userRepo)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
