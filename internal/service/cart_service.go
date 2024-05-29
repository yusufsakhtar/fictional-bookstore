package service

import (
	"fmt"

	"github.com/yusufsakhtar/playstation-assignment/internal/models"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
)

type CartService struct {
	inventoryRepo repository.InventoryRepo
	cartRepo      repository.CartRepo
}

func NewCartService(inventoryRepo repository.InventoryRepo, cartRepo repository.CartRepo) *CartService {
	return &CartService{
		inventoryRepo: inventoryRepo,
		cartRepo:      cartRepo,
	}
}

// TODO: improve both this method and the repo method for adding skus; need better resolution on why items failed to be added
func (s *CartService) AddItemsToUserCart(input repository.AddItemsToUserCartInput) (*repository.AddItemsToUserCartOutput, error) {
	var response repository.AddItemsToUserCartOutput
	for _, sku := range input.SKUs {
		item, err := s.inventoryRepo.GetInventoryItem(repository.GetInventoryItemInput{SKU: sku})
		if err != nil {
			return nil, err
		}
		if item.Stock.Available == 0 {
			response.SKUsFailed = append(response.SKUsFailed, sku)
		} else {
			response.SKUsAdded = append(response.SKUsAdded, sku)
		}
	}
	fmt.Printf("response before adding to cart: %+v\n", response)
	err := s.cartRepo.AddItemsToUserCart(repository.AddItemsToUserCartInput{UserID: input.UserID, SKUs: response.SKUsAdded})
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *CartService) GetCart(input repository.GetCartInput) (*models.Cart, error) {
	return s.cartRepo.GetCart(input)
}

func (s *CartService) ListCarts() ([]*models.Cart, error) {
	return s.cartRepo.ListCarts()
}
