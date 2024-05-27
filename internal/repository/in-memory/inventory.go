package inmemoryrepo

import (
	"encoding/json"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/yusufsakhtar/playstation-assignment/internal/models"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
)

// InMemoryInventoryRepo is a repository that stores inventory data in memory.
type InMemoryInventoryRepo struct {
	inventory map[string]*models.InventoryItem
}

// NewInMemoryInventoryRepo creates a new InMemoryInventoryRepo.
func NewInMemoryInventoryRepo() *InMemoryInventoryRepo {
	return &InMemoryInventoryRepo{
		inventory: make(map[string]*models.InventoryItem),
	}
}

// NewInMemoryInventoryRepoFromFile creates a new InMemoryInventoryRepo from a file for quick bootstrapping.
func NewInMemoryInventoryRepoFromFile(filename string) *InMemoryInventoryRepo {
	file, err := os.Open(filename)
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
}

// CreateInventory creates a new inventory item.
func (r *InMemoryInventoryRepo) CreateInventoryItem(input repository.CreateInventoryItemInput) error {
	// TODO: use something more appropriate than a UUID.
	sku := uuid.New().String()
	r.inventory[sku] = &models.InventoryItem{
		Item: &models.Item{
			SKU:         sku,
			DisplayName: input.DisplayName,
			Price:       input.Price,
		},
		Stock: &models.InventoryItemStock{
			Total:       input.Stock,
			Available:   input.Stock,
			PendingSale: 0,
		},
	}
	return nil
}

// GetInventoryItem retrieves an inventory item by SKU.
func (r *InMemoryInventoryRepo) GetInventoryItem(input repository.GetInventoryItemInput) (*models.Item, error) {
	item, ok := r.inventory[input.SKU]
	if !ok {
		return nil, repository.ErrItemNotFound
	}
	return item.Item, nil
}

// GetInventoryItemStock retrieves the stock of an inventory item by SKU.
func (r *InMemoryInventoryRepo) GetInventoryItemStock(input repository.GetInventoryItemStockInput) (*models.InventoryItemStock, error) {
	item, ok := r.inventory[input.SKU]
	if !ok {
		return nil, repository.ErrItemNotFound
	}
	return item.Stock, nil
}

// DeleteInventoryItem deletes an inventory item by SKU.
func (r *InMemoryInventoryRepo) DeleteInventoryItem(input repository.DeleteInventoryItemInput) error {
	delete(r.inventory, input.SKU)
	return nil
}

// ListItems lists all inventory items.
func (r *InMemoryInventoryRepo) ListItems() ([]*models.Item, error) {
	items := make([]*models.Item, 0, len(r.inventory))
	for _, item := range r.inventory {
		items = append(items, item.Item)
	}
	return items, nil
}
