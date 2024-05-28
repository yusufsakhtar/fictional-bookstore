package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
)

func ListInventory(inventoryRepo repository.InventoryRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := inventoryRepo.ListInventoryItems()
		if err != nil {
			http.Error(w, "Unable to list inventory", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
}

func CreateInventoryItem(inventoryRepo repository.InventoryRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createInventoryItemInput repository.CreateInventoryItemInput
		if err := json.NewDecoder(r.Body).Decode(&createInventoryItemInput); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		if err := inventoryRepo.CreateInventoryItem(createInventoryItemInput); err != nil {
			http.Error(w, "Unable to create inventory item", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func GetInventoryItem(inventoryRepo repository.InventoryRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		sku := vars["sku"]
		item, err := inventoryRepo.GetInventoryItem(repository.GetInventoryItemInput{SKU: sku})
		if err != nil || item == nil {
			http.Error(w, "Unable to retrieve inventory item", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
	}
}

func DeleteInventoryItem(inventoryRepo repository.InventoryRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		sku := vars["sku"]
		if err := inventoryRepo.DeleteInventoryItem(repository.DeleteInventoryItemInput{SKU: sku}); err != nil {
			http.Error(w, "Unable to delete inventory item", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func UpdateInventoryItem(inventoryRepo repository.InventoryRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		sku := vars["sku"]
		var updateInventoryItemInput repository.UpdateInventoryItemInput
		if err := json.NewDecoder(r.Body).Decode(&updateInventoryItemInput); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		updateInventoryItemInput.SKU = sku
		if err := inventoryRepo.UpdateInventoryItem(updateInventoryItemInput); err != nil {
			http.Error(w, "Unable to update inventory item", http.StatusInternalServerError)
			fmt.Printf("UpdateInventoryItem Error: %v\n", err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func UpdateInventoryItemStock(inventoryRepo repository.InventoryRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		sku := vars["sku"]
		var updateInventoryItemStockInput repository.UpdateInventoryItemStockInput
		if err := json.NewDecoder(r.Body).Decode(&updateInventoryItemStockInput); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		updateInventoryItemStockInput.SKU = sku
		if err := inventoryRepo.UpdateInventoryItemStock(updateInventoryItemStockInput); err != nil {
			http.Error(w, "Unable to update inventory item stock", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
