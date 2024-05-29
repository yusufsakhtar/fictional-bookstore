package inmemoryrepo

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/yusufsakhtar/playstation-assignment/internal/models"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
)

// TODO: need an Inventory Service to match pattern elsewhere
type InMemoryInventoryRepo struct {
	inventory map[string]*models.InventoryItem
	mu        sync.RWMutex
}

// NewInMemoryInventoryRepo creates a new InMemoryInventoryRepo, optionally seeding it from a file
func NewInMemoryInventoryRepo(seedFromFile bool, seedFileName string) *InMemoryInventoryRepo {
	if seedFromFile {
		file, err := os.Open(seedFileName)
		if err != nil {
			log.Fatalf("could not open file: %v", err)
		}
		defer file.Close()

		items := make(map[string]*models.InventoryItem)
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&items)
		if err != nil {
			log.Fatalf("could not decode file: %v", err)
		}

		return &InMemoryInventoryRepo{
			inventory: items,
		}
	} else {
		return &InMemoryInventoryRepo{
			inventory: make(map[string]*models.InventoryItem),
		}
	}
}

// CreateInventory creates a new inventory item.
func (r *InMemoryInventoryRepo) CreateInventoryItem(input repository.CreateInventoryItemInput) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// TODO: use something more appropriate than a UUID.
	sku := uuid.New().String()
	r.inventory[sku] = &models.InventoryItem{
		Item: &models.Item{
			SKU:         sku,
			DisplayName: input.DisplayName,
			Price:       input.Price,
		},
		Stock: &models.InventoryItemStock{
			Available:   input.Stock,
			PendingSale: 0,
		},
	}
	return nil
}

// GetInventoryItem retrieves an inventory item by SKU.
func (r *InMemoryInventoryRepo) GetInventoryItem(input repository.GetInventoryItemInput) (*models.InventoryItem, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, ok := r.inventory[input.SKU]
	if !ok {
		return nil, repository.ErrItemNotFound
	}
	return item, nil
}

// GetInventoryItemStock retrieves the stock of an inventory item by SKU.
func (r *InMemoryInventoryRepo) GetInventoryItemStock(input repository.GetInventoryItemStockInput) (*models.InventoryItemStock, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, ok := r.inventory[input.SKU]
	if !ok {
		return nil, repository.ErrItemNotFound
	}
	return item.Stock, nil
}

// DeleteInventoryItem deletes an inventory item by SKU.
func (r *InMemoryInventoryRepo) DeleteInventoryItem(input repository.DeleteInventoryItemInput) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	delete(r.inventory, input.SKU)
	return nil
}

// ListItems lists all inventory items.
func (r *InMemoryInventoryRepo) ListInventoryItems() ([]*models.InventoryItem, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	items := make([]*models.InventoryItem, 0, len(r.inventory))
	for _, item := range r.inventory {
		items = append(items, item)
	}
	return items, nil
}

// UpdateInventoryItem updates an inventory item.
func (r *InMemoryInventoryRepo) UpdateInventoryItem(input repository.UpdateInventoryItemInput) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, ok := r.inventory[input.SKU]
	if !ok {
		return repository.ErrItemNotFound
	}

	if input.DisplayName != nil {
		item.Item.DisplayName = *input.DisplayName
	}
	if input.Price != nil {
		item.Item.Price = *input.Price
	}
	if input.Stock != nil {
		item.Stock.Available = *input.Stock
	}

	return nil
}

func (r *InMemoryInventoryRepo) UpdateInventoryItemStock(input repository.UpdateInventoryItemStockInput) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, ok := r.inventory[input.SKU]
	if !ok {
		return repository.ErrItemNotFound
	}

	item.Stock.Available -= input.AvailableConvertingToPendingSale
	if item.Stock.Available < 0 {
		return repository.ErrInsufficientStock
	}
	item.Stock.PendingSale += input.AvailableConvertingToPendingSale
	return nil
}
