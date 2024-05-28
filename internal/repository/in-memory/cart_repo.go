package inmemoryrepo

import (
	"github.com/google/uuid"
	"github.com/yusufsakhtar/playstation-assignment/internal/models"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
)

type InMemoryCartRepo struct {
	cartsByUserId map[string]*models.Cart
	cartsById     map[string]*models.Cart
}

func NewInMemoryCartRepo() *InMemoryCartRepo {
	return &InMemoryCartRepo{
		cartsByUserId: make(map[string]*models.Cart),
		cartsById:     make(map[string]*models.Cart),
	}
}

func (r *InMemoryCartRepo) CreateCart(input repository.CreateCartInput) error {
	if _, ok := r.cartsByUserId[input.UserID]; ok {
		return repository.ErrCartAlreadyExistsForUser
	}

	cart := &models.Cart{
		ID:      uuid.New().String(),
		UserID:  input.UserID,
		ItemIds: []string{},
	}

	r.cartsByUserId[input.UserID] = cart
	r.cartsById[cart.ID] = cart

	return nil
}

func (r *InMemoryCartRepo) GetCart(input repository.GetCartInput) (*models.Cart, error) {
	cart, ok := r.cartsById[input.ID]
	if !ok {
		return nil, repository.ErrCartNotFound
	}
	return cart, nil
}

func (r *InMemoryCartRepo) GetUserCart(input repository.GetUserCartInput) (*models.Cart, error) {
	cart, ok := r.cartsByUserId[input.UserID]
	if !ok {
		return nil, repository.ErrCartNotFound
	}
	return cart, nil
}

func (r *InMemoryCartRepo) AddItemsToUserCart(input repository.AddItemsToUserCartInput) error {
	cart, ok := r.cartsByUserId[input.UserID]
	if !ok {
		return repository.ErrCartNotFound
	}

	cart.ItemIds = append(cart.ItemIds, input.SKUs...)

	return nil
}
