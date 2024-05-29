package inmemoryrepo

import (
	"encoding/json"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/yusufsakhtar/playstation-assignment/internal/models"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
)

type InMemoryCartRepo struct {
	cartsByUserId map[string]*models.Cart
	cartsById     map[string]*models.Cart
	// TODO: update to use mutex
}

func NewInMemoryCartRepo(seedFromFile bool, seedFileName string) *InMemoryCartRepo {
	if seedFromFile {
		file, err := os.Open(seedFileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		var carts []models.Cart
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&carts)
		if err != nil {
			log.Fatal(err)
		}

		cartsById := make(map[string]*models.Cart)
		cartsByUserId := make(map[string]*models.Cart)
		for _, cart := range carts {
			cartsById[cart.ID] = &cart
			cartsByUserId[cart.UserID] = &cart
		}

		return &InMemoryCartRepo{
			cartsByUserId: cartsByUserId,
			cartsById:     cartsById,
		}
	} else {
		return &InMemoryCartRepo{
			cartsByUserId: make(map[string]*models.Cart),
			cartsById:     make(map[string]*models.Cart),
		}
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

func (r *InMemoryCartRepo) ListCarts() ([]*models.Cart, error) {
	carts := []*models.Cart{}
	for _, cart := range r.cartsById {
		carts = append(carts, cart)
	}
	return carts, nil
}
