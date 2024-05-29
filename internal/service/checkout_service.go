package service

import (
	"github.com/yusufsakhtar/playstation-assignment/internal/models"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
)

// could make this (and all services) an interface but keeping it simple since I have no plans to mock this
type CheckoutService struct {
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

func NewCheckoutService(inventoryRepo repository.InventoryRepo, cartRepo repository.CartRepo, orderRepo repository.OrderRepo) *CheckoutService {
	return &CheckoutService{
		inventoryRepo: inventoryRepo,
		cartRepo:      cartRepo,
		orderRepo:     orderRepo,
	}
}

func (s *CheckoutService) CheckoutCart(input CheckoutInput) (*CheckoutOutput, error) {
	cart, err := s.cartRepo.GetCart(repository.GetCartInput{ID: input.ID})
	if err != nil {
		return nil, err
	}

	var total float64
	var failedSKUs []string
	var successSKUs []string

	// Initialy create the order with all items user is attempting to buy
	order, err := s.orderRepo.CreateOrder(repository.CreateOrderInput{UserID: cart.UserID, SKUs: cart.ItemIds})
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
	err = s.orderRepo.UpdateOrder(repository.UpdateOrderInput{ID: order.ID, UserID: order.UserID, ItemIDs: successSKUs, Status: models.OrderStatusPending})
	if err != nil {
		return nil, err
	}

	// you can imagine another service being used here to calculate taxes, discounts, etc. here
	return &CheckoutOutput{Total: total, SKUsFailed: failedSKUs, SKUsSucceeded: successSKUs}, nil
}
