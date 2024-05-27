package repository

import (
	"errors"

	"github.com/yusufsakhtar/playstation-assignment/internal/models"
)

var ErrUserNotFound = errors.New("user not found")

// Using these interfaces to abstract away the details of the data store from the service layer.
// TODO: implement these interfaces using SQLITE as data store.
type UsersRepo interface {
	CreateUser(input CreateUserInput) error
	GetUser(input GetUserInput) (*models.User, error)
	DeleteUser(input DeleteUserInput) error
	ListUsers() ([]*models.User, error)
}

// Most of these methods would be used as part of an admin console in the product.
// Admin users would manage the inventory.
type InventoryRepo interface {
	CreateItem(input CreateItemInput) error
	GetItem(input GetItemInput) (*models.Item, error)
	DeleteItem(input DeleteItemInput) error
	ListItems() ([]*models.Item, error)
}

type GetUserInput struct {
	ID string `json:"id"`
}

type CreateUserInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type DeleteUserInput struct {
	ID string `json:"id"`
}

type CreateItemInput struct {
	DisplayName string  `json:"display_name"`
	Price       float64 `json:"price"`
}

type GetItemInput struct {
	SKU string `json:"sku"`
}

type DeleteItemInput struct {
	SKU string `json:"sku"`
}
