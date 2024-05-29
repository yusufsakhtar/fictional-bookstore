package service

import (
	"github.com/yusufsakhtar/playstation-assignment/internal/models"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
)

type OrderService struct {
	orderRepo     repository.OrderRepo
	inventoryRepo repository.InventoryRepo
	cartRepo      repository.CartRepo
}

func NewOrderService(
	orderRepo repository.OrderRepo,
	inventoryRepo repository.InventoryRepo,
	cartRepo repository.CartRepo,
) *OrderService {
	return &OrderService{
		orderRepo:     orderRepo,
		inventoryRepo: inventoryRepo,
		cartRepo:      cartRepo,
	}
}

func (s *OrderService) ListOrders() ([]*models.Order, error) {
	return s.orderRepo.ListOrders()
}

func (s *OrderService) CreateOrder(input repository.CreateOrderInput) (*models.Order, error) {
	order, err := s.orderRepo.GetOrderByCartID(repository.GetOrderByCartIDInput{CartID: input.CartID})
	if err == nil {
		return order, repository.ErrOrderAlreadyExistsForCart
	}

	return s.orderRepo.CreateOrder(input)
}

func (s *OrderService) UpdateOrder(input repository.UpdateOrderInput) error {
	return s.orderRepo.UpdateOrder(input)
}

func (s *OrderService) GetOrder(input repository.GetOrderInput) (*models.Order, error) {
	return s.orderRepo.GetOrder(input)
}

func (s *OrderService) ConfirmOrder(input repository.ConfirmOrderInput) error {
	order, err := s.orderRepo.GetOrder(repository.GetOrderInput{ID: input.ID})
	if err != nil {
		return err
	}
	if order.Status != models.OrderStatusPending {
		return repository.ErrInvalidOrderState
	}

	// Future: This could be executed as part of an async process to fulfill a completed order
	for _, sku := range order.ItemIDs {
		err := s.inventoryRepo.UpdateInventoryItemStock(repository.UpdateInventoryItemStockInput{
			SKU:                              sku,
			AvailableConvertingToPendingSale: 1,
		})
		// Throwing an error here bc I haven't build support for re-calculating the total from the order at this point
		// if the item were to become unavailable somehow.
		// A better implementation here might queue up the order for async reprocessing
		// if any items were found to be out of stock/invalid
		if err != nil {
			return err
		}
	}

	// Imagine using order.Total and order.PaymentMethodID to execute a payment hold here
	// and then transacting final payment during the async order fulfillment

	// Update order to completed
	if err = s.orderRepo.UpdateOrder(repository.UpdateOrderInput{
		ID: input.ID,
		CreateOrderInput: repository.CreateOrderInput{
			UserID:  order.UserID,
			CartID:  order.CartID,
			ItemIDs: order.ItemIDs,
			Total:   order.Total,
			Status:  models.OrderStatusCompleted,
		},
	}); err != nil {
		return err
	}

	// Finally, delete the cart object
	if err = s.cartRepo.DeleteCart(repository.DeleteCartInput{ID: order.CartID}); err != nil {
		return err
	}
	return nil
}
