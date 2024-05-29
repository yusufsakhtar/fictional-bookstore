package service

import (
	"github.com/yusufsakhtar/playstation-assignment/internal/models"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
)

type OrderService struct {
	orderRepo repository.OrderRepo
}

func NewOrderService(orderRepo repository.OrderRepo) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
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
