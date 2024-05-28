package service

import "github.com/yusufsakhtar/playstation-assignment/internal/repository"

// could make this (and all services) an interface but keeping it simple since I have no plans to mock this
type CheckoutService struct {
	inventoryRepo repository.InventoryRepo
	cartRepo      repository.CartRepo
}
