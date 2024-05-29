package service

import (
	"fmt"

	"github.com/yusufsakhtar/playstation-assignment/internal/models"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
)

type CartService struct {
	inventoryRepo repository.InventoryRepo
	cartRepo      repository.CartRepo
	orderRepo     repository.OrderRepo
}

type CheckoutInput struct {
	ID string
}

type CheckoutOutput struct {
	Total         float64
	SKUsFailed    []string
	SKUsSucceeded []string
}

func NewCartService(inventoryRepo repository.InventoryRepo, cartRepo repository.CartRepo, orderRepo repository.OrderRepo) *CartService {
	return &CartService{
		inventoryRepo: inventoryRepo,
		cartRepo:      cartRepo,
		orderRepo:     orderRepo,
	}
}

func (s *CartService) GetCart(input repository.GetCartInput) (*models.Cart, error) {
	return s.cartRepo.GetCart(input)
}

func (s *CartService) ListCarts() ([]*models.Cart, error) {
	return s.cartRepo.ListCarts()
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

func (s *CartService) CheckoutCart(input CheckoutInput) (*CheckoutOutput, error) {
	cart, err := s.cartRepo.GetCart(repository.GetCartInput{ID: input.ID})
	if err != nil {
		return nil, err
	}

	var total float64
	var failedSKUs []string
	var successSKUs []string

	// Initialy create the order with all items user is attempting to buy
	// Note: this could be a good candidate for a message on a queue in a distributed system
	// TODO: change total on this struct to be a pointer to allow null values
	order, err := s.orderRepo.CreateOrder(repository.CreateOrderInput{
		UserID:  cart.UserID,
		CartID:  cart.ID,
		ItemIDs: cart.ItemIds,
		Total:   total,
		Status:  models.OrderStatusCreated,
	})
	if err != nil {
		return nil, err
	}

	for _, sku := range cart.ItemIds {
		// I'm realizing this method is not as useful as expected since we make a call to get
		// the item details in the succeeding method.
		// TODO: simplify/cleanup
		err := s.inventoryRepo.UpdateInventoryItemStock(repository.UpdateInventoryItemStockInput{SKU: sku, AvailableConvertingToPendingSale: 1})
		if err != nil {
			if err == repository.ErrInsufficientStock {
				failedSKUs = append(failedSKUs, sku)
			} else {
				return nil, err
			}
		} else {
			item, err := s.inventoryRepo.GetInventoryItem(repository.GetInventoryItemInput{SKU: sku})
			if err != nil {
				return nil, err
			}
			total += item.Item.Price
			successSKUs = append(successSKUs, sku)
		}
	}
	// Update order with specific items that were successfully purchased and pending status
	// Note: this could be a good candidate for a message on a queue in a distributed system
	err = s.orderRepo.UpdateOrder(repository.UpdateOrderInput{
		ID: order.ID,
		CreateOrderInput: repository.CreateOrderInput{
			UserID:  order.UserID,
			CartID:  order.CartID,
			ItemIDs: successSKUs,
			Total:   total,
			Status:  models.OrderStatusPending,
		},
	})
	if err != nil {
		return nil, err
	}

	// you can imagine another service being used here to calculate taxes, discounts, etc. here
	return &CheckoutOutput{Total: total, SKUsFailed: failedSKUs, SKUsSucceeded: successSKUs}, nil
}
