package repository

import (
	"errors"

	"github.com/yusufsakhtar/playstation-assignment/internal/models"
)

// TODO: move these to a pkg/errors
var ErrUserNotFound = errors.New("user not found")
var ErrItemNotFound = errors.New("item not found")
var ErrInsufficientStock = errors.New("insufficient stock")
var ErrCartNotFound = errors.New("cart not found")
var ErrCartAlreadyExistsForUser = errors.New("cart already exists for given user")
var ErrOrderNotFound = errors.New("order not found")
var ErrOrderAlreadyExistsForCart = errors.New("order already exists for given cart")
var ErrInvalidOrderState = errors.New("invalid order state")

// Using these interfaces to abstract away the details of the data store from the service layer.
// In a next iteration, we could implement these interfaces using SQLITE as the data store.
type UserRepo interface {
	CreateUser(input CreateUserInput) (*models.User, error)
	GetUser(input GetUserInput) (*models.User, error)
	DeleteUser(input DeleteUserInput) error
	ListUsers() ([]*models.User, error)
}

// Most of these methods would be used as part of an admin console in the product.
// Admin users would manage the inventory.
type InventoryRepo interface {
	CreateInventoryItem(input CreateInventoryItemInput) error
	GetInventoryItem(input GetInventoryItemInput) (*models.InventoryItem, error)
	GetInventoryItemStock(input GetInventoryItemStockInput) (*models.InventoryItemStock, error)
	DeleteInventoryItem(input DeleteInventoryItemInput) error
	ListInventoryItems() ([]*models.InventoryItem, error)
	UpdateInventoryItem(input UpdateInventoryItemInput) error
	UpdateInventoryItemStock(input UpdateInventoryItemStockInput) error
}

type CartRepo interface {
	CreateCart(input CreateCartInput) error
	GetCart(input GetCartInput) (*models.Cart, error)
	ListCarts() ([]*models.Cart, error)
	GetUserCart(input GetUserCartInput) (*models.Cart, error)
	AddItemsToUserCart(input AddItemsToUserCartInput) error
	DeleteCart(input DeleteCartInput) error
}

type OrderRepo interface {
	CreateOrder(input CreateOrderInput) (*models.Order, error)
	UpdateOrder(input UpdateOrderInput) error
	ListOrders() ([]*models.Order, error)
	GetOrder(input GetOrderInput) (*models.Order, error)
	GetOrderByCartID(input GetOrderByCartIDInput) (*models.Order, error)
}

// TODO: implement this as part of extending the checkout/order flow to account for payment
type PaymentRepo interface {
	CreatePaymentMethod(input CreatePaymentMethodInput) error
	GetPaymentMethod(input GetPaymentMethodInput) (*models.PaymentMethod, error)
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

type CreateInventoryItemInput struct {
	DisplayName string  `json:"display_name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

type GetInventoryItemStockInput struct {
	SKU string `json:"sku"`
}

type GetInventoryItemInput struct {
	SKU string `json:"sku"`
}

type DeleteInventoryItemInput struct {
	SKU string `json:"sku"`
}

type UpdateInventoryItemInput struct {
	SKU         string   `json:"sku"`
	DisplayName *string  `json:"display_name,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Stock       *int     `json:"stock,omitempty"`
}

// Not super happy with the interface here. The goal was to provide a way to update stock without needing to set each of the total/available/pending stock explicitly - Im sure there's a better way to do this but sticking with this for now.
type UpdateInventoryItemStockInput struct {
	SKU                              string `json:"sku"`
	AvailableConvertingToPendingSale int    `json:"available_converting_to_pending_sale"`
}

type CreateCartInput struct {
	UserID string `json:"user_id"`
}

type GetCartInput struct {
	ID string `json:"id"`
}

type GetUserCartInput struct {
	UserID string `json:"user_id"`
}

type AddItemsToUserCartInput struct {
	UserID string   `json:"user_id"`
	SKUs   []string `json:"skus"`
}

type AddItemsToUserCartOutput struct {
	SKUsAdded  []string `json:"skus_added"`
	SKUsFailed []string `json:"skus_failed"`
}

type CreateOrderInput struct {
	UserID  string             `json:"user_id"`
	CartID  string             `json:"cart_id"`
	ItemIDs []string           `json:"skus"`
	Total   float64            `json:"total"`
	Status  models.OrderStatus `json:"status"`
}
type UpdateOrderInput struct {
	ID string `json:"id"`
	CreateOrderInput
}

type CreatePaymentMethodInput struct {
	Type    models.PaymentMethodType `json:"type"`
	Balance float64                  `json:"balance"`
}

type GetPaymentMethodInput struct {
	ID string `json:"id"`
}

type GetOrderInput struct {
	ID string `json:"id"`
}

type GetOrderByCartIDInput struct {
	CartID string `json:"cart_id"`
}

type ConfirmOrderInput struct {
	ID string `json:"id"`
}

type DeleteCartInput struct {
	ID string `json:"id"`
}
