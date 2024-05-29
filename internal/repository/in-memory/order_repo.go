package inmemoryrepo

import (
	"github.com/google/uuid"
	"github.com/yusufsakhtar/playstation-assignment/internal/models"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
)

// TODO: add orderService to handle business logic, let repo simply handle data access
type InMemoryOrderRepo struct {
	ordersByUserId map[string][]*models.Order
	ordersById     map[string]*models.Order
	ordersByCartId map[string]*models.Order
	// TODO: update to use mutex
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{
		ordersByUserId: make(map[string][]*models.Order),
		ordersById:     make(map[string]*models.Order),
		ordersByCartId: make(map[string]*models.Order),
	}
}

func (r *InMemoryOrderRepo) CreateOrder(input repository.CreateOrderInput) (*models.Order, error) {
	if _, ok := r.ordersByUserId[input.UserID]; !ok {
		r.ordersByUserId[input.UserID] = []*models.Order{}
	}

	order := &models.Order{
		ID:      uuid.New().String(),
		UserID:  input.UserID,
		CartID:  input.CartID,
		ItemIDs: input.ItemIDs,
		Total:   input.Total,
		Status:  models.OrderStatusCreated,
	}

	r.ordersByUserId[input.UserID] = append(r.ordersByUserId[input.UserID], order)
	r.ordersById[order.ID] = order
	r.ordersByCartId[input.CartID] = order

	return order, nil
}

func (r *InMemoryOrderRepo) UpdateOrder(input repository.UpdateOrderInput) error {
	order, ok := r.ordersById[input.ID]
	if !ok {
		return repository.ErrOrderNotFound
	}

	order.ItemIDs = input.ItemIDs
	order.Status = input.Status
	order.UserID = input.UserID
	return nil
}

func (r *InMemoryOrderRepo) ListOrders() ([]*models.Order, error) {
	orders := []*models.Order{}
	for _, order := range r.ordersById {
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *InMemoryOrderRepo) GetOrder(input repository.GetOrderInput) (*models.Order, error) {
	order, ok := r.ordersById[input.ID]
	if !ok {
		return nil, repository.ErrOrderNotFound
	}
	return order, nil
}

func (r *InMemoryOrderRepo) GetOrderByCartID(input repository.GetOrderByCartIDInput) (*models.Order, error) {
	order, ok := r.ordersByCartId[input.CartID]
	if !ok {
		return nil, repository.ErrOrderNotFound
	}
	return order, nil
}
