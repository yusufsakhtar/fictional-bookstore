package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
)

func GetOrder(orderRepo repository.OrderRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		order, err := orderRepo.GetOrder(repository.GetOrderInput{ID: id})
		if err != nil {
			if errors.Is(err, repository.ErrOrderNotFound) {
				http.Error(w, "order not found", http.StatusNotFound)
				return
			}
			http.Error(w, "internal server error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(order)
	}
}

func ListOrders(orderRepo repository.OrderRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orders, err := orderRepo.ListOrders()
		if err != nil {
			http.Error(w, "internal server error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(orders)
	}
}
